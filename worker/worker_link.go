package worker

import (
	pb "com.github.patrickz98.omnet/proto"
	"com.github.patrickz98.omnet/simple"
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"runtime"
	"time"
)

func Link(config Config) (err error) {

	logger.Println("config", simple.PrettyString(config))

	//
	// Set up a connection to the server
	//

	conn, err := grpc.Dial(config.BrokerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.Fatalf("did not connect: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	client := pb.NewBrokerClient(conn)

	md := metadata.New(map[string]string{
		"workerId": config.WorkerId,
	})

	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, md)

	link, err := client.Link(ctx)
	if err != nil {
		logger.Fatalln(err)
	}

	wClient := workerClient{
		config:        config,
		link:          link,
		freeResources: runtime.NumCPU(),
	}

	go func() {
		for {
			tasks, err := link.Recv()
			if err != nil {
				logger.Println(err)
				return
			}

			err = wClient.OccupyResource()
			if err != nil {
				logger.Println(err)
				return
			}

			go func(tasks *pb.Tasks) {
				byt, err := json.MarshalIndent(tasks, "", "  ")
				if err != nil {
					logger.Println(err)
					return
				}

				randWait := rand.Intn(300) + 30
				waitTine := time.Duration(randWait) * time.Millisecond
				logger.Printf("doing %s for %v\n", byt, waitTine)
				time.Sleep(waitTine)

				err = wClient.FeeResource()
				if err != nil {
					logger.Println(err)
					return
				}
			}(tasks)
		}
	}()

	for {
		err = wClient.SendClientInfo()
		if err != nil {
			break
		}

		time.Sleep(time.Second * 23)
	}

	return
}
