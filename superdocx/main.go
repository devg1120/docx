package main


import    (
	"superdocx/libs"
)

func main() {
	//Open file
	doc,err := superdocx.Open("./Example.docx")
	if err != nil {
		panic(err)
	}
	//Create and append paragraph
	para := doc.Document.Body.CreateAddParagraph()
	//Create and append run in paragraph
	run := para.CreateRun()
	//set style bold
	run.Style.SetBold(true)
	//set text of run
	run.SetText("Hello World")
	para1 := doc.Document.Body.CreateAddParagraph()
	para1.CreateRuns(2)
	//size is in half-points (half-poits = fontsize*2)
	para1.Runs[0].Style.SetSize(20)
	para1.Runs[0].Style.SetItalic(true)
	para1.Runs[0].SetText("run0 ")
	para1.Runs[1].Style.SetSize(21)
	para1.Runs[1].SetText("run1 ")
	para1.Runs[2].Style.SetSize(21)
	para1.Runs[2].Style.SetCaps(true)
	para1.Runs[2].SetText("run2")
	//replace text even if it is across runs. new text using same style of the fist old text letter
	para1.Replace("run0 run1 ru", "new text to replace old with italic style")
	err = doc.SaveAs("./NewDocument.docx")
	// or ([]byte ,error) doc.SaveToMemory()
	if err != nil {
		panic(err)
	}
}
