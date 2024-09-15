# [ plot ] Package

## Overview

The `plot` package is a wrapper around the powerful [github.com/go-echarts/go-echarts](https://github.com/go-echarts/go-echarts) library, designed to simplify data visualization. It provides an easy-to-use interface for generating common chart types, such as bar charts, while also giving users the flexibility to leverage the full power of `go-echarts` for more advanced customizations.

This package is perfect for users who want to quickly visualize their data without needing to write verbose code. Advanced users can still access and configure all the underlying `go-echarts` features for more complex use cases.

---

## Installation

```bash
go get github.com/HazelnutParadise/insyra/plot
go get github.com/go-echarts/go-echarts/v2
go get github.com/go-rod/rod/lib/proto
go get github.com/luabagg/orcgen/v2
```

---

## Usage Example

### 1. Using `map[string][]float64`

This example shows how to create a bar chart using `map[string][]float64` as the `SeriesData`.

```go
package main

import (
	"github.com/HazelnutParadise/insyra/plot"
)

func main() {
	config := plot.BarChartConfig{
		Title:      "Sales Data",
		Subtitle:   "Monthly Sales",
		XAxis:      []string{"January", "February", "March"},
		SeriesData: map[string][]float64{
			"Product A": {120, 200, 150},
			"Product B": {80, 160, 90},
		},
		ShowLabels: true,
		Colors:     []string{"#5470C6", "#91CC75"},
	}

	// Create a bar chart
	barChart := plot.CreateBarChart(config)

	// Save the chart as an HTML file
	plot.SaveHTML(barChart, "sales_data_map.html")

	// Save the chart as a PNG file
	plot.SavePNG(barChart, "sales_data_map.png")
}
```

### 2. Using `insyra.DataList`

This example demonstrates how to create a bar chart using `[]*insyra.DataList` as the `SeriesData`.

```go
package main

import (
	"github.com/HazelnutParadise/insyra"
	"github.com/HazelnutParadise/insyra/plot"
)

func main() {
	// Create DataList objects for different products
	dataListA := insyra.NewDataList(120, 200, 150).SetName("Product A")

	dataListB := insyra.NewDataList(80, 160, 90).SetName("Product B")

	config := plot.BarChartConfig{
		Title:      "Sales Data",
		Subtitle:   "Monthly Sales",
		XAxis:      []string{"January", "February", "March"},
		SeriesData: []*insyra.DataList{dataListA, dataListB},
		ShowLabels: true,
		Colors:     []string{"#5470C6", "#91CC75"},
	}

	// Create a bar chart
	barChart := plot.CreateBarChart(config)

	// Save the chart as an HTML file
	plot.SaveHTML(barChart, "sales_data_datalist.html")

	// Save the chart as a PNG file
	plot.SavePNG(barChart, "sales_data_datalist.png")
}
```

---

## Features

### Bar Chart

#### `BarChartConfig`

- `Title`: The title of the chart.
- `Subtitle`: The subtitle of the chart.
- `XAxis`: Data for the X-axis (categories).
- `SeriesData`: The data for the series. Supported types:
  - `map[string][]float64`: A map where keys are series names, and values are data points.
  - `[]*insyra.DataList`: A list of `DataList` structures.
  - `[]insyra.IDataList`: A list of `IDataList` interface implementations.
- `XAxisName` (optional): Name for the X-axis.
- `YAxisName` (optional): Name for the Y-axis.
- `YAxisNameGap` (optional): Gap between the Y-axis name and the chart.
- `Colors` (optional): Colors for the bars.
- `ShowLabels` (optional): Display labels on the bars.
- `LabelPos` (optional): Position of the labels (e.g., "top", "bottom"). Default is "top".
- `GridTop` (optional): Space between the top of the chart and the title. Default is `"80"`.

#### `CreateBarChart`

`func CreateBarChart(config BarChartConfig) *charts.Bar`

Creates a bar chart based on the provided `BarChartConfig` and returns a `*charts.Bar` object, which can be customized further using `go-echarts` options.

### Saving Charts

#### `SaveHTML`

`func SaveHTML(chart Renderable, path string)`

Renders the chart and saves it as an HTML file at the specified path.

#### `SavePNG`

`func SavePNG(chart Renderable, pngPath string)`

Renders the chart as a PNG file and saves it to the specified path. This utilizes the `orcgen` library to convert the chart into a PNG.

---

## Advanced Customization

While the `plot` package simplifies the process of creating charts, it also provides full access to the underlying `go-echarts` chart objects. This allows advanced users to configure and extend the charts beyond the default settings offered by `plot`.

For instance, after creating a bar chart using `CreateBarChart`, users can call `go-echarts` methods to further customize the chart:

```go
barChart.SetGlobalOptions(
	charts.WithTitleOpts(opts.Title{
		Title:    "Custom Title",
		Subtitle: "Custom Subtitle",
	}),
	charts.WithXAxisOpts(opts.XAxis{
		Name: "Custom X-Axis Name",
	}),
)
```

This flexibility ensures that users can start with simple visualizations and evolve them into more complex representations using `go-echarts`.

---

## Error Handling

- **Unsupported `SeriesData` types**: If an unsupported data type is passed to `SeriesData`, a warning will be logged and the chart will not be created.
- **File and rendering errors**: Both `SaveHTML` and `SavePNG` will log fatal errors if they encounter issues while rendering or saving the files.