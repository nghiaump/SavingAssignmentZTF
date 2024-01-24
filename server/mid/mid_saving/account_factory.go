package main

// ISavingFactory used Factory Pattern
type ISavingFactory interface {
	CreateSavingAccountPT(termType string, term int32) *SavingAccountPT
}

type SavingFactoryKYC2 struct{}
type SavingFactoryKYC3 struct{}

func (factoryKYC2 *SavingFactoryKYC2) CreateSavingAccountPT(termType string, term int32) *SavingAccountPT {
	var rate float32
	switch termType {
	case "DAYS":
		{
			if term == 21 {
				rate = 0.03
			}
		}
	case "MONTHS":
		{
			if term == 3 {
				rate = 0.04
			} else if term == 6 {
				rate = 0.05
			} else if term == 12 {
				rate = 0.06
			}
		}
	}
	return &SavingAccountPT{InterestRate: rate}
}

func (factoryKYC3 *SavingFactoryKYC3) CreateSavingAccountPT(termType string, term int32) *SavingAccountPT {
	var rate float32
	switch termType {
	case "DAYS":
		{
			if term == 21 {
				rate = 0.035
			}
		}
	case "MONTHS":
		{
			if term == 3 {
				rate = 0.045
			} else if term == 6 {
				rate = 0.055
			} else if term == 12 {
				rate = 0.065
			}
		}
	}
	return &SavingAccountPT{InterestRate: rate}
}
