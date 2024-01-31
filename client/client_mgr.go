package client

import (
	"errors"

	"github.com/kehiy/RoboPac/log"
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
)

type Mgr struct {
	clients map[string]IClient
}

func NewClientMgr() *Mgr {
	return &Mgr{
		clients: make(map[string]IClient),
	}
}

func (cm *Mgr) AddClient(addr string, c IClient) {
	cm.clients[addr] = c
}

func (cm *Mgr) GetRandomClient() IClient {
	for _, c := range cm.clients {
		return c
	}

	return nil
}

func (cm *Mgr) GetBlockchainInfo() (*pactus.GetBlockchainInfoResponse, error) {
	for _, c := range cm.clients {
		info, err := c.GetBlockchainInfo()
		if err != nil {
			continue
		}
		return info, nil
	}

	return nil, errors.New("unable to get blockchain info")
}

func (cm *Mgr) GetBlockchainHeight() (uint32, error) {
	for _, c := range cm.clients {
		height, err := c.GetBlockchainHeight()
		if err != nil {
			continue
		}
		return height, nil
	}

	return 0, errors.New("unable to get blockchain height")
}

func (cm *Mgr) GetLastBlockTime() (uint32, uint32) {
	var lastBlockTime uint32 = 0
	var lastBlockHeight uint32 = 0
	for _, c := range cm.clients {
		t, h, err := c.LastBlockTime()
		if err != nil {
			continue
		}
		if t > lastBlockTime {
			lastBlockTime = t
			lastBlockHeight = h
		}
	}

	return lastBlockTime, lastBlockHeight
}

func (cm *Mgr) GetNetworkInfo() (*pactus.GetNetworkInfoResponse, error) {
	for name, c := range cm.clients {
		if name == "local-node" {
			continue
		}

		info, err := c.GetNetworkInfo()
		if err != nil {
			continue
		}
		return info, nil
	}

	return nil, errors.New("unable to get network info")
}

func (cm *Mgr) GetPeerInfoFirstVal(address string) (*pactus.PeerInfo, error) {
	for name, c := range cm.clients {
		if name == "local-node" {
			continue
		}

		networkInfo, err := c.GetNetworkInfo()
		if err != nil {
			continue
		}

		if networkInfo != nil {
			for _, p := range networkInfo.ConnectedPeers {
				for i, addr := range p.ConsensusAddress {
					if addr == address {
						if i != 0 {
							return nil, errors.New("please enter the first validator address")
						}
						return p, nil
					}
				}
			}
		}
	}

	return nil, errors.New("peer does not exist")
}

func (cm *Mgr) GetPeerInfo(address string) (*pactus.PeerInfo, error) {
	for name, c := range cm.clients {
		if name == "local-node" {
			continue
		}

		networkInfo, err := c.GetNetworkInfo()
		if err != nil {
			continue
		}

		if networkInfo != nil {
			for _, p := range networkInfo.ConnectedPeers {
				for _, addr := range p.ConsensusAddress {
					if addr == address {
						return p, nil
					}
				}
			}
		}
	}

	return nil, errors.New("peer does not exist")
}

func (cm *Mgr) IsStakedValidator(address string) bool {
	for _, c := range cm.clients {
		val, err := c.GetValidatorInfo(address)
		if err != nil {
			log.Info("error", "err", err)
			return false
		}
		log.Info("passed", "bool", val.Validator.Stake > 0)
		return val.Validator.Stake > 0
	}

	log.Info("passed")
	return false
}

func (cm *Mgr) GetValidatorInfo(address string) (*pactus.GetValidatorResponse, error) {
	for _, c := range cm.clients {
		val, err := c.GetValidatorInfo(address)
		if err != nil {
			continue
		}
		return val, nil
	}

	return nil, errors.New("unable to get validator info")
}

func (cm *Mgr) GetValidatorInfoByNumber(num int32) (*pactus.GetValidatorResponse, error) {
	for _, c := range cm.clients {
		val, err := c.GetValidatorInfoByNumber(num)
		if err != nil {
			continue
		}
		return val, nil
	}

	return nil, errors.New("unable to get validator info")
}

func (cm *Mgr) GetTransactionData(txID string) (*pactus.GetTransactionResponse, error) {
	for _, c := range cm.clients {
		txData, err := c.GetTransactionData(txID)
		if err != nil {
			continue
		}
		return txData, nil
	}

	return nil, errors.New("unable to get transaction data")
}

func (cm *Mgr) Stop() {
	for addr, c := range cm.clients {
		if err := c.Close(); err != nil {
			log.Error("could not close connection to RPC node", "err", err, "RPCAddr", addr)
		}
	}
}
