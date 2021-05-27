package notionapi

const (
	ObjectTypeDatabase    ObjectType = "database"
	ObjectTypeTitle       ObjectType = "title"
	ObjectTypeText        ObjectType = "text"
	ObjectTypeRichText    ObjectType = "rich_text"
	ObjectTypeCheckbox    ObjectType = "checkbox"
	ObjectTypeSelect      ObjectType = "select"
	ObjectTypeNumber      ObjectType = "number"
	ObjectTypeFormula     ObjectType = "formula"
	ObjectTypeDate        ObjectType = "date"
	ObjectTypeRelation    ObjectType = "relation"
	ObjectTypeRollup      ObjectType = "rollup"
	ObjectTypeMultiSelect ObjectType = "multi_select"
	ObjectTypePeople      ObjectType = "people"
	ObjectTypeFiles       ObjectType = "files"
	ObjectTypeList        ObjectType = "list"
	ObjectTypeHeading1    ObjectType = "heading_1"
	ObjectTypeHeading2    ObjectType = "heading_2"
	ObjectTypeHeading3    ObjectType = "heading_3"
	ObjectTypeParagraph   ObjectType = "paragraph"
	ObjectTypeToggle      ObjectType = "toggle"
	ObjectTypeUser        ObjectType = "user"
	ObjectTypePerson      ObjectType = "person"
	ObjectTypeBot         ObjectType = "bot"

	ObjectTypeBulletedListItem ObjectType = "bulleted_list_item"
	ObjectTypeNumberedListItem ObjectType = "numbered_list_item"

	ObjectTypeToDo        ObjectType = "to_do"
	ObjectTypeChildPage   ObjectType = "child_page"
	ObjectTypeUnsupported ObjectType = "unsupported"
)

type ObjectType string

func (ot ObjectType) String() string {
	return string(ot)
}

type ObjectID string

func (oID ObjectID) String() string {
	return string(oID)
}

type BasicObject struct {
	ID          ObjectID           `json:"id"`
	Type        ObjectType         `json:"type"`
	Title       *TextObject        `json:"title,omitempty"`
	Text        *TextObject        `json:"text,omitempty"`
	RichText    *RichTextObject    `json:"rich_text,omitempty"`
	Checkbox    *struct{}          `json:"checkbox,omitempty"`
	Select      *SelectObject      `json:"select,omitempty"`
	Number      *NumberObject      `json:"number,omitempty"`
	Formula     *FormulaObject     `json:"formula,omitempty"`
	Date        *struct{}          `json:"date,omitempty"`
	Relation    *RelationObject    `json:"relation,omitempty"`
	Rollup      *RollupObject      `json:"rollup,omitempty"`
	MultiSelect *MultiSelectObject `json:"multi_select,omitempty"`
	People      *struct{}          `json:"people,omitempty"`
	Files       *struct{}          `json:"files,omitempty"`
	Paragraph   *ParagraphObject   `json:"paragraph,omitempty"`
	Toggle      *Toggle            `json:"toggle,omitempty"`
}

type Color string

func (c Color) String() string {
	return string(c)
}

type RichTextObject struct {
	Type        ObjectType  `json:"type"`
	Text        TextObject  `json:"text"`
	Annotations Annotations `json:"annotations"`
	PlainText   string      `json:"plain_text"`
	Href        string      `json:"href"`
}

type TextObject struct {
	Content string `json:"content"`
	Link    string `json:"link"`
}

type Annotations struct {
	Bold          bool  `json:"bold"`
	Italic        bool  `json:"italic"`
	Strikethrough bool  `json:"strikethrough"`
	Underline     bool  `json:"underline"`
	Code          bool  `json:"code"`
	Color         Color `json:"color"`
}

type ParagraphObject []RichTextObject

type SelectObject struct {
	Options []SelectObject `json:"options"`
}

type SelectOption struct {
	ID    ObjectID
	Name  string `json:"name"`
	Color Color  `json:"color"`
}

type FormatType string

func (ft FormatType) String() string {
	return string(ft)
}

type NumberObject struct {
	Format FormatType `json:"format"`
}

type FormulaObject struct {
	Value string `json:"value"`
}

type RelationObject struct {
	Database           DatabaseID `json:"database"`
	SyncedPropertyName string     `json:"synced_property_name"`
}

type FunctionType string

func (ft FunctionType) String() string {
	return string(ft)
}

type RollupObject struct {
	RollupPropertyName   string       `json:"rollup_property_name"`
	RelationPropertyName string       `json:"relation_property_name"`
	RollupPropertyID     ObjectID     `json:"rollup_property_id"`
	RelationPropertyID   ObjectID     `json:"relation_property_id"`
	Function             FunctionType `json:"function"`
}

type MultiSelectObject struct {
	Options []SelectOption `json:"options"`
}

type Toggle struct {
	Text RichTextObject `json:"text"`
}

type Cursor string

func (c Cursor) String() string {
	return string(c)
}
