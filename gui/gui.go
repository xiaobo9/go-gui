package gui

import (
	"fmt"

	. "github.com/lxn/walk/declarative"
)

func NewDataModel() *DataModel {
	model := new(DataModel)
	model.items = make([]*Data, 3)

	model.items[0] = &Data{Index: 0, Name: "a", Price: 20}
	model.items[1] = &Data{Index: 1, Name: "b", Price: 18}
	model.items[2] = &Data{Index: 2, Name: "c", Price: 19}

	return model
}

func (mw *MyDataMainWindow) show() {
	Window()
}

func Window() {
	mw := &MyDataMainWindow{model: NewDataModel()}
	window := MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Data展示",
		Icon:     "favicon.ico",
		Size:     Size{Width: 800, Height: 600},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					HSpacer{},
					PushButton{Text: "Add", OnClicked: func() { addBtn(mw) }},
					PushButton{Text: "Delete", OnClicked: func() { deleteBtn(mw) }},
					PushButton{Text: "ExecChecked", OnClicked: func() { execChecked(mw) }},
					PushButton{Text: "AddPriceChecked", OnClicked: func() { addPriceChecked(mw) }},
				},
			},
			Composite{
				Layout: VBox{},
				ContextMenuItems: []MenuItem{
					Action{Text: "I&nfo", OnTriggered: mw.tvItemActivated},
					Action{Text: "E&xit", OnTriggered: func() { mw.Close() }},
				},
				Children: []Widget{
					TableView{
						AssignTo:         &mw.tableView,
						CheckBoxes:       true,
						ColumnsOrderable: true,
						MultiSelection:   true,
						Columns: []TableViewColumn{
							{Title: "编号"},
							{Title: "名称"},
							{Title: "价格"},
						},
						Model:                 mw.model,
						OnCurrentIndexChanged: func() { currentIndexChanged(mw) },
						OnItemActivated:       mw.tvItemActivated,
					},
				},
			},
		},
	}
	_, e := window.Run()
	if e != nil {
		fmt.Println(e)
	}
}

func addBtn(mw *MyDataMainWindow) {
	mw.model.items = append(mw.model.items, &Data{
		Index: mw.model.Len() + 1,
		Name:  "啥名字",
		Price: mw.model.Len() * 5,
	})
	mw.model.PublishRowsReset()
	mw.tableView.SetSelectedIndexes([]int{})
}

func deleteBtn(mw *MyDataMainWindow) {
	items := []*Data{}
	remove := mw.tableView.SelectedIndexes()
	for i, item := range mw.model.items {
		removeOk := false
		for _, j := range remove {
			if i == j {
				removeOk = true
				break
			}
		}
		if !removeOk {
			items = append(items, item)
		}
	}
	mw.model.items = items
	mw.model.PublishRowsReset()
	mw.tableView.SetSelectedIndexes([]int{})
}

func execChecked(mw *MyDataMainWindow) {
	for _, item := range mw.model.items {
		if item.checked {
			fmt.Printf("checked: %v\n", item)
		}
	}
}

func addPriceChecked(mw *MyDataMainWindow) {
	for i, item := range mw.model.items {
		if item.checked {
			item.Price++
			mw.model.PublishRowChanged(i)
		}
	}
}

func currentIndexChanged(mw *MyDataMainWindow) {
	i := mw.tableView.CurrentIndex()
	if 0 <= i {
		fmt.Printf("OnCurrentIndexChanged: %v\n", mw.model.items[i].Name)
	}
}
