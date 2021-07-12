package consumer

import (
	"context"
	pb "github.com/pzierahn/project.go.omnetpp/proto"
	"log"
)

func (pConn *providerConnection) init(cons *consumer) (err error) {

	simulation := cons.simulation

	source := &checkoutObject{
		SimulationId: simulation.Id,
		Filename:     "source.tgz",
		Data:         cons.simulationSource,
	}

	if err = pConn.checkout(source); err != nil {
		return
	}

	if err = pConn.setupExecutable(simulation); err != nil {
		return
	}

	stream, err := pConn.provider.Allocate(context.Background())
	if err != nil {
		// TODO: Do not crash here
		log.Fatalln(err)
	}

	go func() {

		log.Printf("[%s] startAllocator", pConn.name())

		for {
			alloc, err := stream.Recv()
			if err != nil {
				break
			}

			log.Printf("[%s] allocated %d slots",
				pConn.name(), alloc.Slots)

			for inx := uint32(0); inx < alloc.Slots; inx++ {

				task, ok := <-cons.allocator
				if !ok {
					//
					// No tasks left
					//

					return
				}

				go func() {
					// TODO: Find a better way to handle this
					defer cons.finished.Done()

					err := pConn.run(task)
					if err != nil {
						// TODO: Don't crash here!
						log.Fatalln(pConn.name(), err)
					}
				}()
			}
		}
	}()

	go func() {
		//
		// Communicate changes in the allocSlots number to the provider
		//

		cond := cons.allocCond

		cond.L.Lock()
		allocateJobs := uint32(len(cons.allocate))
		cond.L.Unlock()

		for {
			log.Printf("[%s] request %d slots", pConn.name(), allocateJobs)

			err := stream.Send(&pb.AllocateRequest{
				SimulationId: simulation.Id,
				Request:      allocateJobs,
			})
			if err != nil {
				log.Println(err)
				break
			}

			cond.L.Lock()
			cond.Wait()
			allocateJobs = uint32(len(cons.allocate))
			cond.L.Unlock()
		}
	}()

	return
}
