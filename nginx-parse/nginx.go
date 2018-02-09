package nxparse

// NginxConf Nginx 配置信息
type NginxConf struct {
	*Section
}

// Section 配置段
type Section struct {
	data       map[string]Line
	childs     []*Section
	childsKey  []string
	preSection *Section
}

// Line 配置行
type Line []string

// NewSection 申请 Section 的内存空间
func NewSection() *Section {
	sec := &Section{}
	sec.data = make(map[string]Line)
	return sec
}

// NewChildSection 生成子段
func (sec *Section) NewChildSection(key string) *Section {
	child := &Section{}
	child.data = make(map[string]Line)
	child.preSection = sec
	sec.AppendChilds(key, child)
	return child
}

// AppendLine 段增加行
func (sec *Section) AppendLine(k string, l Line) {
	sec.data[k] = append(sec.data[k], l...)
}

// AppendChilds 增加子段
func (sec *Section) AppendChilds(k string, v *Section) {
	sec.childs = append(sec.childs, v)
	sec.childsKey = append(sec.childsKey, k)
}

// GetChilds 获取子段
func (sec *Section) GetChilds(k string) []*Section {
	ret := make([]*Section, 0)
	for i, key := range sec.childsKey {
		if k == key {
			ret = append(ret, sec.childs[i])
		}
	}
	return ret
}

// GetFistChild 获取子段
func (sec *Section) GetFistChild(k string) *Section {
	for i, key := range sec.childsKey {
		if k == key {
			return sec.childs[i]
		}
	}
	return nil
}

// GetValues 获取值数组
func (sec *Section) GetValues(k string) Line {
	return sec.data[k]
}

// GetValue 获取值
func (sec *Section) GetValue(k string) string {
	if len(sec.data[k]) > 0 {
		return sec.data[k][0]
	}
	return ""
}

// Init 初始化读取 Nginx 配置
func Init(filename string) (*NginxConf, error) {
	ret := &NginxConf{}
	ret.Section = NewSection()

	err := parseSection(ret.Section, filename)
	return ret, err
}
