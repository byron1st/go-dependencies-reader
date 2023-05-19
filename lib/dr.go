package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

type DependencyRelation struct {
	Caller string `json:"caller"`
	Callee string `json:"callee"`
}

func (dr DependencyRelation) Print() error {
	b, err := json.Marshal(dr)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, string(b))
	return nil
}
