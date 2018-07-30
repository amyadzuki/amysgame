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

vec3 HueToRgb(float hue) {
	float r = abs(hue * 6 - 3) - 1;
	float g = 2 - abs(hue * 6 - 2);
	float b = 2 - abs(hue * 6 - 4);
	return clamp(vec3(r, g, b), 0, 1);
}

vec3 HslToRgb(vec3 hsl) {
	vec3 rgb = HueToRgb(hsl.x);
	float c = (1 - abs(2 * hsl.z - 1)) * hsl.y;
	return clamp((rgb - 0.5) * c + hsl.z, 0, 1);
}

vec3 HsvToRgb(vec3 hsv) {
	vec3 k = vec3(3, 2, 1) * (1.0/3);
	vec3 p = abs(fract(hsv.xxx + k.xyz) * 6.0 - vec3(3, 3, 3));
	return hsv.b * mix(k.xxx, clamp(p - k.xxx, 0, 1), hsv.y);
}

void main() {
#if MAT_TEXTURES>0
	vec4 sampColor, color;
	vec3 hslSkin, rgbSkin;
	sampColor = texture(MatTexture[0], vTexcoord);
	hslSkin = sampColor.rgb;
	hslSkin += HsmSkinDelta.xyz;
	hslSkin.r -= floor(hslSkin.r);
	if (hslSkin.r < 0) {
		hslSkin.r += 1;
	}
	rgbSkin = HsvToRgb(hslSkin);
	color = vec4(rgbSkin.rgb, 1);
#if MAT_TEXTURES>1
	vec4 uwfc, uw;
	uwfc = texture(MatTexture[1], vTexcoord);
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
