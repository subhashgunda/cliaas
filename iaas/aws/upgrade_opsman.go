package aws

import (
	"fmt"
	"time"

	errwrap "github.com/pkg/errors"
)

type UpgradeOpsMan struct {
	client ClientAPI
}

func NewUpgradeOpsMan(configs ...func(*UpgradeOpsMan) error) (*UpgradeOpsMan, error) {
	upgradeOpsMan := new(UpgradeOpsMan)
	for _, cfg := range configs {
		err := cfg(upgradeOpsMan)
		if err != nil {
			return nil, errwrap.Wrap(err, "new upgradeOpsMan config loading error")
		}
	}

	if upgradeOpsMan.client == nil {
		return nil, errwrap.New("must configure client")
	}
	return upgradeOpsMan, nil
}

func ConfigClient(value ClientAPI) func(*UpgradeOpsMan) error {
	return func(upgradeOpsMan *UpgradeOpsMan) error {
		upgradeOpsMan.client = value
		return nil
	}
}

func (s *UpgradeOpsMan) Upgrade(name, vpc, ami, instanceType, ip string) error {
	instance, err := s.client.GetVMInfo(fmt.Sprintf("%s*", name))
	if err != nil {
		return err
	}
	err = s.client.StopVM(*instance)
	if err != nil {
		return err
	}
	t := time.Now()
	dateString := fmt.Sprintf("%d-%02d-%02dT%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	newName := fmt.Sprintf("%s - %s", name, dateString)
	newInstance, err := s.client.CreateVM(*instance, ami, instanceType, newName)
	if err != nil {
		return err
	}
	err = s.client.WaitForStartedVM(newName)
	if err != nil {
		return err
	}
	err = s.client.AssignPublicIP(*newInstance, ip)
	if err != nil {
		return err
	}
	return s.client.DeleteVM(*instance)
}
