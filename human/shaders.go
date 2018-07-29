package human

var HumanSkinVs = `
#include <attributes>
// blank line required by preprocessor
uniform mat4 MVP;
out vec2 vTexcoord;
void main() {
	vec4 position;
	position = vec4(VertexPosition.xyz, 1);
	gl_Position = MVP * position;
	vTexcoord = VertexTexcoord.xy;
}
`

var HumanSkinFs = `
#include <material>
// blank line required by preprocessor
in vec2 vTexcoord;
out vec4 fColor;
void main() {
#if MAT_TEXTURES>0
	fColor = vec4(texture(MatTexture[0], vTexcoord));
#if MAT_TEXTURES>1
	// TODO: underwear
#endif
#else
	fColor = vec4(1, 0, 1, 1);
#endif
}
`
