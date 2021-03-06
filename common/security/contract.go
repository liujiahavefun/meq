package security

import (
	"errors"
)

// The contract's state possible values.
const (
	ContractStateAllowed = uint8(iota)
	ContractStateRefused
)

// Contract represents an interface for a contract.
type Contract interface {
	Validate(key Key) bool // Validate checks the security key with the contract.
}

// contract represents a contract (user account).
type contract struct {
	ID        uint32 `json:"id"`     // Gets or sets the contract id.
	MasterID  uint16 `json:"master"` // Gets or sets the master id.
	Signature uint32 `json:"sign"`   // Gets or sets the signature of the contract.
	State     uint8  `json:"state"`  // Gets or sets the state of the contract.
}

// Validate validates the contract data against a key.
func (c *contract) Validate(key Key) bool {
	return c.MasterID == key.Master() &&
		c.Signature == key.Signature() &&
		c.ID == key.Contract() &&
		c.State == ContractStateAllowed
}

// ContractProvider represents an interface for a contract provider.
// type ContractProvider interface {
// 	Create() (Contract, error)
// 	Get(id uint32) (Contract, bool)
// }

// SingleContractProvider provides contracts on premise.
type SingleContractProvider struct {
	owner *contract // The owner contract.
}

// NewSingleContractProvider creates a new single contract provider.
func NewSingleContractProvider(license *License) *SingleContractProvider {
	p := new(SingleContractProvider)
	p.owner = new(contract)
	p.owner.MasterID = 1
	p.owner.ID = license.Contract
	p.owner.Signature = license.Signature
	return p
}

// Name returns the name of the provider.
func (p *SingleContractProvider) Name() string {
	return "single"
}

// Configure configures the provider.
func (p *SingleContractProvider) Configure(config map[string]interface{}) error {
	return nil
}

// Create creates a contract, the SingleContractProvider way.
func (p *SingleContractProvider) Create() (Contract, error) {
	return nil, errors.New("Single contract provider can not create contracts")
}

// Get returns a ContractData fetched by its id.
func (p *SingleContractProvider) Get(id uint32) (Contract, bool) {
	if p.owner == nil || p.owner.ID != id {
		return nil, false
	}

	return p.owner, true
}
