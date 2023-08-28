package algs

import (
	"sort"

	"github.com/mjevans93308/platformscience/pkg/models"
	"github.com/mjevans93308/platformscience/util/helpers"
	"github.com/mjevans93308/platformscience/util/localctx"
)

type NameInfo struct {
	name string
	numC int
	numV int
}

// CalculateSS take the ctx and the file contents of names and addresses
// We will need to compute a SS for each name -> address pairing, meaning a n^2 runtime is probably unavoidable
// To help with this, we will precalculate each name's number of vowels and consonants
// After calculating each SS, we will order the list by greatest SS to least, so that our highest SS are at the head
// We will then iterate through that list to populate a new slice, starting with the greatest SS
// and checking that we aren't re-using any names
// This final list will be our output
func CalculateSS(ctx *localctx.Localctx, content *models.FileContent) []models.NameAddressSS {
	var nameObjs []NameInfo
	for _, name := range content.Names {
		var nameObj NameInfo
		cons, vows := helpers.CountChars(name)
		nameObj.numC = cons
		nameObj.numV = vows
		nameObj.name = name
		nameObjs = append(nameObjs, nameObj)
	}

	var nassl []models.NameAddressSS
	for _, nameObj := range nameObjs {
		for _, addr := range content.Addresses {
			ss := 1.0
			if len(addr)%2 == 0 {
				ss = 1.5 * float64(nameObj.numV)
			} else {
				ss = float64(nameObj.numC)
			}
			if helpers.ShareCommonFactors(len(nameObj.name), len(addr)) {
				ss += ss * 0.5
			}
			nass := models.NameAddressSS{
				Name:    nameObj.name,
				Address: addr,
				SS:      ss,
			}
			nassl = append(nassl, nass)
		}
	}

	// sort by SS to get best results at the front
	sort.Slice(nassl, func(i, j int) bool {
		return nassl[i].SS > nassl[j].SS
	})

	// our nassl list should be ordered with the greatest SS vals at the top
	// starting with those, insert them into a map so that we can build a list
	// of individual addr and name pairings with the greatest SS
	// once we insert a name, make sure we don't reuse that name later
	resMap := make(map[string]bool)
	namesMap := make(map[string]bool)
	var resL []models.NameAddressSS
	for _, nass := range nassl {
		_, resFound := resMap[nass.Address]
		_, nameFound := namesMap[nass.Name]
		if !resFound && !nameFound {
			resMap[nass.Address] = true
			namesMap[nass.Name] = true
			resL = append(resL, nass)
		}
	}

	return resL
}
