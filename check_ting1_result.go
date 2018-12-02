package main

import (
	"fmt"
	"strings"
	"github.com/fatih/color"
)

type checkTing1Result struct {
	avgImproveNum   float64
	improveWayCount int
	avgTingCount    float64
}

// TODO: 提醒切这张牌可以断幺！
// TODO: 赤牌改良提醒！！
func (r *checkTing1Result) Print() {
	if r.improveWayCount > 0 {
		if r.improveWayCount >= 100 {
			fmt.Printf("%-6.2f[%3d改良]", r.avgImproveNum, r.improveWayCount)
		} else {
			fmt.Printf("%-6.2f[%2d 改良]", r.avgImproveNum, r.improveWayCount)
		}
	} else {
		fmt.Print(strings.Repeat(" ", 15))
	}

	fmt.Print(" ")
	color.New(getTingCountColor(r.avgTingCount)).Printf("%5.2f", r.avgTingCount)
	fmt.Print(" 听牌数")

	fmt.Println()
}