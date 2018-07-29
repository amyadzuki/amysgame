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
	vTexcoord = vec2(VertexTexcoord.x, 1.0 - VertexTexcoord.y);
}
`

var HumanSkinFs = `
#include <material>
// blank line required by preprocessor
in vec2 vTexcoord;
out vec4 fColor;
void main() {
#if MAT_TEXTURES>0
	fColor = texture(MatTexture[0], vTexcoord);
#if MAT_TEXTURES>1
	vec4 uwfc, uw;
	uwfc = texture(MatTexture[1], vTexcoord);
	uw = mix(vec4(0, 0, 0, 0), vec4(1, 1, 1, 1), uwfc.r);
	uw = mix(uw, vec4(0.875, 0.875, 0.875, 0.5), uwfc.g);
	uw = mix(uw, vec4(0xff/255.0, 0xb6/255.0, 0xc1/255.0, 1), uwfc.b);
	fColor = mix(fColor, uw, uwfc.a);
#endif
#else
	fColor = vec4(1, 0, 1, 1);
#endif
}
`
