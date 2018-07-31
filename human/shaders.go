package human

import (
	"strconv"
	"unsafe"
)

// Skin ///////////////////////////////////////////////////////////////////////

var HumanSkinVs = `
#include <attributes>
// blank line required by preprocessor
uniform mat4 MVP;
out vec2 vTexcoord;
void main() {
	vec4 position;
	position = vec4(VertexPosition.xyz, 1);
	gl_Position = MVP * position;
	vTexcoord = vec2(VertexTexcoord.x, 1.0 - VertexTexcoord.y); // TODO: flip textures
}
`

var HumanSkinFs = `
#include <material>
// blank line required by preprocessor
in vec2 vTexcoord;
uniform vec4 HumanSkin[` + strconv.Itoa(int(unsafe.Sizeof(HumanSkinMaterialUdata{}) / 16)) + `];
#define HumanSkinDelta HumanSkin[0]
#define HumanUwFabric HumanSkin[1]
#define HumanUwDetail HumanSkin[2]
#define HumanUwTrim HumanSkin[3]
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
	hsvSkin = mix(light.xyz, dark.xyz, HumanSkinDelta.w * HumanSkinDelta.w);
#else
	vec4 sampColor;
	sampColor = texture(MatTexture[0], vTexcoord);
	hsvSkin = sampColor.xyz;
#endif
	float hue = hsvSkin.x;
	hsvSkin *= HumanSkinDelta.xyz * 2;
	hsvSkin.x = fract(hue + HumanSkinDelta.x + 0.5);
	rgbSkin = HsvToRgb(hsvSkin);
	color = vec4(rgbSkin.rgb, 1);
#if MAT_TEXTURES>2
	vec4 uwfc, uw;
	uwfc = texture(MatTexture[2], vTexcoord);
	uw = mix(color, HumanUwFabric, uwfc.r);
	uw = mix(uw, HumanUwDetail, uwfc.g);
	uw = mix(uw, HumanUwTrim, uwfc.b);
	color = mix(color, uw, uwfc.a);
#endif
#else
	color = vec4(1, 0, 1, 1);
#endif
	fColor = vec4(color.rgb, 1);
}
`

// Eyes ///////////////////////////////////////////////////////////////////////

var HumanEyesVs = `
#include <attributes>
// blank line required by preprocessor
uniform mat4 MVP;
out vec4 vPosition;
out vec2 vTexcoord;
void main() {
	vec4 position = MVP * vec4(VertexPosition.xyz, 1);
	gl_Position = position;
	vPosition = position;
	vTexcoord = vec2(VertexTexcoord.x, 1.0 - VertexTexcoord.y); // TODO: flip textures
}
`

var HumanEyesFs = `
#include <material>
// blank line required by preprocessor
in vec4 vPosition;
in vec2 vTexcoord;
uniform vec4 HumanEyes[4];
#define HumanEyesColor HumanEyes[0]
out vec4 fColor;

void main() {
    vec4 color;
#if MAT_TEXTURES>0
    vec4 sampColor = texture(MatTexture[0], vTexcoord);
    vec4 invSamp = 1 - sampColor;
    float scalar = invSamp.x * invSamp.y * invSamp.z;
    vec4 mixColor = max(sampColor, vec4(HumanEyesColor.rgb, 1));
    float mixAmt = clamp(vPosition.z / vPosition.w, 0, 1);
    mixAmt = mixAmt * mixAmt; // **2
    mixAmt = mixAmt * mixAmt; // **4
    mixAmt = mixAmt * mixAmt; // **8
    mixAmt = mixAmt * mixAmt; // **16
    mixAmt = mixAmt * mixAmt; // **32
    mixAmt = mixAmt * mixAmt; // **64
    mixAmt = clamp(mixAmt * scalar, 0, 0.5);
    color = mix(sampColor, mixColor, mixAmt);
    color = vec4(mix(HumanEyesColor.rgb, color.rgb, color.a), 1);
    if (min(vTexcoord.x, vTexcoord.y) >= 0.8125) {
        discard;
        // color = sampColor.rgba; // <-- this doesn't seem to work right
    }
#else
    color = vec3(1, 0, 1, 1);
#endif
    fColor = color;
}
`
