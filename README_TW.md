# Insyra

[![Test](https://github.com/HazelnutParadise/insyra/actions/workflows/test.yml/badge.svg)](https://github.com/HazelnutParadise/insyra/actions/workflows/test.yml)
[![Go version](https://img.shields.io/github/go-mod/go-version/HazelnutParadise/insyra.svg)](https://github.com/HazelnutParadise/insyra)
[![Go Report Card](https://goreportcard.com/badge/github.com/HazelnutParadise/insyra)](https://goreportcard.com/report/github.com/HazelnutParadise/insyra)
[![GoDoc](https://godoc.org/github.com/HazelnutParadise/insyra?status.svg)](https://pkg.go.dev/github.com/HazelnutParadise/insyra)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

Go 語言次世代資料分析庫。

![logo](logo/logo.webp)

## 太快、太美、太簡單。

**Insyra** 庫是一個動態且多功能的 Go 語言資料分析工具。提供了豐富的功能集，可用於數據操作、統計計算、資料視覺化等，對於處理複雜數據結構的開發者來說，是一個必不可少的工具包。

> [!NOTE]
> 如果文檔中的某些功能無法使用，可能是該功能還未包含在最新發布的版本中。請至 **[Releases](https://github.com/HazelnutParadise/insyra/releases)** 查看對應版本源碼中的文檔。

> [!IMPORTANT]
> **對於 Insyra 文檔中未明確列出的任何函數或方法，表示該功能仍在積極開發中。這些實驗性功能可能會提供不穩定的結果。**    
>
> 請參閱我們 **[文檔](https://github.com/HazelnutParadise/insyra/tree/main/Docs)** 資料夾中的最新更新以獲取更多詳細資訊。

## 開始使用

### 安裝

要開始使用 **Insyra**，請使用以下命令進行安裝：

```sh
go get github.com/HazelnutParadise/insyra
```

### 快速範例

```go
package main

import (
    "fmt"
    "github.com/HazelnutParadise/insyra"
)

func main() {
    dl := insyra.NewDataList(1, 2, 3, 4, 5)
    dl.Append(6)
    fmt.Println("DataList:", dl.Data())
    fmt.Println("Mean:", dl.Mean())
}
```

## DataList

`DataList` 是 **Insyra** 的核心結構，能夠存儲、管理和分析動態數據集合。它提供了各種用於數據操作和統計分析的方法。

有關方法和功能的完整列表，請參閱 **[DataList 文檔](https://github.com/HazelnutParadise/insyra/tree/main/Docs/DataList.md)**。

## DataTable

`DataTable` 結構提供了表格數據的表示方式，允許以結構化格式存儲和操作數據。它提供了數據過濾、排序和聚合的方法，使其成為數據分析的強大工具。

**您還可以僅用一行代碼在 DataTables 和 CSV 文件之間進行轉換，實現與外部數據源的無縫整合。**

有關方法和功能的完整列表，請參閱 **[DataTable 文檔](https://github.com/HazelnutParadise/insyra/tree/main/Docs/DataTable.md)**。

## 套件

**Insyra** 還提供了多個擴展套件，每個都專注於數據分析的特定方面。

### **[stats](/Docs/stats.md)**

提供數據分析的統計函數，包括偏度、峰度和矩計算。

### **[parallel](/Docs/parallel.md)**

為數據操作和分析提供平行處理能力。可用於執行任何函數，並自動等待所有 goroutine 完成。

### **[plot](/Docs/plot.md)**

強大的 [github.com/go-echarts/go-echarts](https://github.com/go-echarts/go-echarts) 庫的封裝，用於簡化資料視覺化。

### **[gplot](/Docs/gplot.md)**

基於 [github.com/gonum/plot](https://github.com/gonum/plot) 的視覺化套件。快速且不需要 Chrome。

### **[lpgen](/Docs/lpgen.md)**

提供一個 **超級簡單** 且直觀的方式來生成線性規劃（LP）模型並將其保存為 `.lp` 檔。支援設置目標、添加約束、定義變量邊界，並指定二進制或整數變量。

### **[lp](/Docs/lp.md)**

使用 [GLPK](https://www.gnu.org/software/glpk/) 的全自動線性規劃（LP）包。

## 進階使用

除了基本用法外，**Insyra** 還提供了處理不同數據類型和執行複雜統計操作的強大功能。請在 **[詳細文檔](https://github.com/HazelnutParadise/insyra/tree/main/Docs)** 中探索更多內容。

## 貢獻

歡迎各種形式的貢獻！您可以通過以下方式貢獻 **Insyra**：
- **Issues：** 提出問題、建議或功能請求。
- **Pull Requests：** 提交代碼更改或新功能。
- **Discussions：** 參與討論，分享您的想法和建議。

<!-- 有關詳細信息，請參閱 [貢獻指南](https://github.com/HazelnutParadise/insyra/blob/main/CONTRIBUTING.md)。 -->

## 貢獻者
![contributors](https://contrib.rocks/image?repo=HazelnutParadise/insyra)

## 授權

Insyra 採用 MIT 許可證授權。請參閱 [LICENSE](https://github.com/HazelnutParadise/insyra/blob/main/LICENSE) 文件以獲取更多資訊。
