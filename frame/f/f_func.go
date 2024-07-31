package f

import (
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/frame/frouter"
)

func App() *fapp.App { return fapp.Now() }

func Router() frouter.Router { return fapp.Now().Router }
