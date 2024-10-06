package initialize

import "BuildService/client/receiver"

type Clients struct {
	ReceiverClient receiver.IReceiverClient
}

func NewClients() *Clients {
	receiverClient := receiver.NewReceiverClient()
	return &Clients{
		ReceiverClient: receiverClient,
	}
}