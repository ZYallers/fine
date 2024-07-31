package consts

const TemplateGenServiceContentHead = `
// ================================================================================
// Code generated and maintained by Fine CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// {CreatedAt}.
// ================================================================================

package {PackageName}

{Imports}
`

const TemplateGenServiceContentInterface = `
{InterfaceName} interface {
	{FuncDefinition}
}
`

const TemplateGenServiceContentVariable = `
local{StructName} {InterfaceName}
`

const TemplateGenServiceContentRegister = `
func {StructName}() {InterfaceName} {
	if local{StructName} == nil {
		panic("implement not found for interface {InterfaceName}, forgot register?")
	}
	return local{StructName}
}

func Register{StructName}(i {InterfaceName}) {
	local{StructName} = i
}
`

const TemplateGenServiceLogicContent = `
// ==========================================================================
// Code generated and maintained by Fine CLI tool. DO NOT EDIT.
// {CreatedAt}.
// ==========================================================================

package {PackageName}

import(
	{Imports}
)
`
