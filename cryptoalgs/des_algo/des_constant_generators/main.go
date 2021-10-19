package main

import (
    "text/template"
    "strconv"
    "strings"
    "os"

    "github.com/antchfx/htmlquery"
    "golang.org/x/net/html"
)

const (
    IPPermutationAmount = 64
    ExpandEFunctionAmount = 48

    SBlockAmount = 8
    SRowsAmount = 4
    SColumnsAmount = 16

    PPermutationAmount = 32
    KeyPermutationAmount = 56
    KeyShiftAmount = 16
    KeyPostPermutationAmount = 48

    ConstantTemplateFilePath = "des_constants.template"
    ConstantTemplateSavePath = "../des/des_constants.go"
)

type TemplateData struct {
    IPPermutationData       []byte
    ExpandEFunctionData     []byte
    SPermutationData        [][][]byte
    PPermutationData        []byte
    KeyPermutationData      []byte
    KeyShiftData            []byte
    KeyPostPermutationData  []byte
    IPPostPermutationData   []byte
}

func main() {

    doc, err := htmlquery.LoadURL("https://ru.wikipedia.org/wiki/DES")

    if err != nil {
        panic(err)
    }

    // Generate template data struct
    templateData := TemplateData{
        IPPermutationData:      parseOneDimensionArray(
            doc,
            `//caption[contains(string(.), "Таблица 1.")]/following-sibling::tbody//td`,
            IPPermutationAmount,
        ),
        ExpandEFunctionData:    parseOneDimensionArray(
            doc,
            `//caption[contains(string(.), "Таблица 2.")]/following-sibling::tbody//td`,
            ExpandEFunctionAmount,
        ),
        SPermutationData:       parseThreeDimensionArray(
            doc,
            `//caption[contains(string(.), "Таблица 3.")]/following-sibling::tbody/tr[position(.) > 2]/th[position(.) > 1 and not(@style)]`,
            SBlockAmount, SRowsAmount, SColumnsAmount,
        ),
        PPermutationData:       parseOneDimensionArray(
            doc,
            `//caption[contains(string(.), "Таблица 4.")]/following-sibling::tbody//td`,
            PPermutationAmount,
        ),
        KeyPermutationData:     parseOneDimensionArray(
            doc,
            `//caption[contains(string(.), "Таблица 5.")]/following-sibling::tbody//td[not(@style)]`,
            KeyPermutationAmount,
        ),
        KeyShiftData:           parseOneDimensionArray(
            doc,
            `//caption[contains(string(.), "Таблица 6.")]/following-sibling::tbody//td[position(.) > 1]`,
            KeyShiftAmount,
        ),
        KeyPostPermutationData: parseOneDimensionArray(
            doc,
            `//caption[contains(string(.), "Таблица 7.")]/following-sibling::tbody//td`,
            KeyPostPermutationAmount,
        ),
        IPPostPermutationData:  parseOneDimensionArray(
            doc,
            `//caption[contains(string(.), "Таблица 8.")]/following-sibling::tbody//td`,
            IPPermutationAmount,
        ),
    }

    outputTmpl := template.Must(template.New(ConstantTemplateFilePath).Funcs(map[string]interface{}{
        "join": func(elems []byte) string {
            elemsToJoin := make([]string, 0, len(elems))

            for _, elem := range elems {
                elemsToJoin = append(elemsToJoin, strconv.Itoa(int(elem)))
            }

            return strings.Join(elemsToJoin, ", ")
        },
    }).Delims("<<", ">>").ParseFiles(ConstantTemplateFilePath))

    fd, err := os.OpenFile(ConstantTemplateSavePath, os.O_WRONLY | os.O_CREATE, 0755)

    if err != nil {
        panic(err)
    }

    if err = outputTmpl.Execute(fd, &templateData); err != nil {
        panic(err)
    }
}

func parseOneDimensionArray(doc *html.Node, xpathString string, lenCheck int) ([]byte) {
    res := make([]byte, lenCheck)

    nodes, err := htmlquery.QueryAll(
        doc,
        xpathString,
    )

    if err != nil {
        panic(err)
    }

    if len(nodes) != lenCheck {
        panic("Invalid amount of nodes parsed")
    }

    for i, node := range nodes {
        if intval, err := strconv.ParseInt(strings.ReplaceAll(node.FirstChild.Data, "\n", ""), 10, 8); err != nil {
            panic(err)
        } else {
            res[i] = byte(intval)
        }
    }

    return res
}

func parseThreeDimensionArray(doc *html.Node, xpathString string, zlen, xlen, ylen int) ([][][]byte) {
    res := make([][][]byte, zlen)

    for i := 0; i < zlen; i++ {
        res[i] = make([][]byte, xlen)

        for j := 0; j < xlen; j++ {
            res[i][j] = make([]byte, ylen)
        }
    }

    nodes, err := htmlquery.QueryAll(
        doc,
        xpathString,
    )

    if err != nil {
        panic(err)
    }

    if len(nodes) != zlen * xlen * ylen {
        panic("Invalid amount of values parsed!")
    }

    nodeCounter := 0

    for z := 0; z < zlen; z++ {
        for x := 0; x < xlen; x++ {
            for y := 0; y < ylen; y++ {
                if intval, err := strconv.ParseInt(nodes[nodeCounter].FirstChild.Data, 10, 8); err != nil {
                    panic(err)
                } else {
                    res[z][x][y] = byte(intval)
                    nodeCounter++
                }
            }
        }
    }

    return res
}
