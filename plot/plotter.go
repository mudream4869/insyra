package plot

import (
	"image/color"
	"os"

	"github.com/HazelnutParadise/insyra" // 假設 insyra 是你的主包
)

// Plotter 是畫圖包的主要結構
type Plotter struct {
	data   *insyra.DataList
	title  string
	xLabel string
	yLabel string
	color  color.Color
}

// NewPlotter 創建一個新的 Plotter
func NewPlotter(data *insyra.DataList, title, xLabel, yLabel string, color color.Color) *Plotter {
	return &Plotter{
		data:   data,
		title:  title,
		xLabel: xLabel,
		yLabel: yLabel,
		color:  color,
	}
}

// LinePlot 畫出折線圖
func (p *Plotter) LinePlot(outputFile string) error {
	// 這裡實現折線圖繪製的邏輯
	// 使用第三方包或者原生的 Go 庫來處理圖形繪製
	return nil
}

// BarPlot 畫出柱狀圖
func (p *Plotter) BarPlot(outputFile string) error {
	// 這裡實現柱狀圖繪製的邏輯
	return nil
}

// SavePlot 將圖形保存到文件中
func (p *Plotter) SavePlot(outputFile string) error {
	// 使用原生的 Go 庫或者第三方包來保存圖像
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()
	// 繪製完成的圖像寫入文件中
	return nil
}
