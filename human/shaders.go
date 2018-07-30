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
uniform vec4 HumanSkinMaterial[4];
#define HsmSkinColor HumanSkinMaterial[0]
#define HsmUwFabric HumanSkinMaterial[1]
#define HsmUwDetail HumanSkinMaterial[2]
#define HsmUwTrim HumanSkinMaterial[3]
out vec4 fColor;
void main() {
	vec4 transpSkin;
	vec3 skin;
	transpSkin = texture(MatTexture[0], vTexcoord);
	skin = mix(HsmSkinColor.rgb, transpSkin.rgb, transpSkin.a);
#if MAT_TEXTURES>0
	fColor = vec4(skin.rgb, 1);
#if MAT_TEXTURES>1
	vec4 uwfc, uw;
	uwfc = texture(MatTexture[1], vTexcoord);
	uw = mix(fColor, HsmUwFabric, uwfc.r);
	uw = mix(uw, HsmUwDetail, uwfc.g);
	uw = mix(uw, HsmUwTrim, uwfc.b);
	fColor = mix(fColor, uw, uwfc.a);
#endif
#else
	fColor = vec4(1, 0, 1, 1);
#endif
	fColor.a = 1;
}
`
