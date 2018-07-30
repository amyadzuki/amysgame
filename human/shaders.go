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
uniform vec4 HumanSkin[4];
#define HsmSkinDelta HumanSkin[0]
#define HsmUwFabric HumanSkin[1]
#define HsmUwDetail HumanSkin[2]
#define HsmUwTrim HumanSkin[3]
out vec4 fColor;

vec3 HsvToRgb(vec3 hsv) {
	vec3 k = vec3(3, 2, 1) / 3;
	vec3 p = abs(fract(hsv.xxx + k.xyz) * 6.0 + vec3(-3));
	return hsv.z * mix(k.xxx, clamp(p - k.xxx, 0, 1), hsv.y);
}

void main() {
#if MAT_TEXTURES>0
	vec4 color;
	vec3 hsvSkin, rgbSkin;
#if MAT_TEXTURES>1
	vec4 dark, light;
	dark = texture(MatTexture[0], vTexcoord);
	light = texture(MatTexture[1], vTexcoord);
	hsvSkin = mix(dark.xyz, light.xyz, HsmSkinDelta.w);
#else
	vec4 sampColor;
	sampColor = texture(MatTexture[0], vTexcoord);
	hsvSkin = sampColor.xyz;
#endif
	float hue = hsvSkin.x;
	hsvSkin *= HsmSkinDelta.xyz * 2;
	hsvSkin.x = hue + HsmSkinDelta.x;
	hsvSkin.x -= floor(hsvSkin.x);
	if (hsvSkin.x < 0) {
		hsvSkin.x += 1;
	}
	rgbSkin = HsvToRgb(hsvSkin);
	color = vec4(rgbSkin.rgb, 1);
#if MAT_TEXTURES>2
	vec4 uwfc, uw;
	uwfc = texture(MatTexture[2], vTexcoord);
	uw = mix(color, HsmUwFabric, uwfc.r);
	uw = mix(uw, HsmUwDetail, uwfc.g);
	uw = mix(uw, HsmUwTrim, uwfc.b);
	color = mix(color, uw, uwfc.a);
#endif
#else
	color = vec4(1, 0, 1, 1);
#endif
	fColor = vec4(color.rgb, 1);
}
`
