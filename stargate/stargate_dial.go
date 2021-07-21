package stargate

import (
	"context"
	"github.com/lucas-clemente/quic-go"
	"github.com/pzierahn/project.go.omnetpp/equic"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

func DialUDP(ctx context.Context, dialAddr DialAddr) (conn *net.UDPConn, peer *net.UDPAddr, err error) {

	log.Printf("DialUDP: dialAddr=%v", dialAddr)

	conn, err = net.ListenUDP("udp", &net.UDPAddr{})
	if err != nil {
		return
	}

	if deadline, ok := ctx.Deadline(); ok {
		log.Printf("DialUDP: deadline=%v", deadline)
		err = conn.SetDeadline(deadline)
		if err != nil {
			return
		}

		defer func() {
			// Reset deadline
			_ = conn.SetDeadline(time.Time{})
		}()
	}

	client := stargateClient{
		conn: conn,
		dial: dialAddr,
	}
	peer, err = client.resolvePeer()
	if err != nil {
		return
	}

	//
	// TODO: Start a go-routine after the peer was resolved to prevent connection jamming
	//

	log.Printf("DialUDP: dialAddr=%v peer=%v", dialAddr, peer)

	var wg sync.WaitGroup
	var once sync.Once

	cctx, cnl := context.WithTimeout(ctx, time.Second*5)
	defer cnl()

	if deadline, ok := cctx.Deadline(); ok {
		err = conn.SetReadDeadline(deadline)
		if err != nil {
			return
		}

		defer func() {
			// Reset deadline
			_ = conn.SetDeadline(time.Time{})
		}()
	}

	helper := dialer{
		conn: conn,
		peer: peer,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		sendErr := helper.sendHellos(cctx)
		if sendErr != nil {
			once.Do(func() {
				err = sendErr
			})
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		receiveErr := helper.receive()
		if receiveErr != nil {
			once.Do(func() {
				err = receiveErr
			})
		}
	}()

	wg.Wait()

	if err != nil {
		return
	}

	// Everything is okay
	log.Printf("DialUDP: dialAddr=%v peer=%v connection established", dialAddr, peer)

	return
}

func DialGRPCClientConn(ctx context.Context, dialAddr DialAddr) (conn *grpc.ClientConn, err error) {
	gate, remote, err := DialUDP(ctx, dialAddr)
	if err != nil {
		// Connection failed!
		return
	}

	conn, err = grpc.Dial(
		remote.String(),
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithContextDialer(equic.GRPCDialer(gate)),
	)

	return
}

func DialQUICListener(ctx context.Context, dialAddr DialAddr) (p2p quic.Listener, err error) {

	conn, _, err := DialUDP(ctx, dialAddr)
	if err != nil {
		return
	}

	tlsConf, _ := equic.GenerateTLSConfig()

	return quic.Listen(conn, tlsConf, equic.Config)
}

func DialQUICgRPCListener(ctx context.Context, dialAddr DialAddr) (p2p net.Listener, err error) {

	qLis, err := DialQUICListener(ctx, dialAddr)
	if err != nil {
		return
	}

	return equic.Listen(qLis), err
}
