#version 330
uniform mat4 MVP;
uniform mat4 M;
uniform mat4 V;

uniform vec3 lightPosition_worldSpace;

in vec3 position_modelSpace;

out vec3 vertexNormal_cameraSpace;
out vec3 lightDirection_cameraSpace;

void main() {
    gl_Position = MVP * vec4(position_modelSpace, 1);

    vec3 position_worldSpace = (M * vec4(position_modelSpace, 1)).xyz;
    vec3 position_cameraSpace = (V * M * vec4(position_modelSpace, 1)).xyz;

    vec3 eyeDirection_cameraSpace = vec3(0, 0, 0) - position_cameraSpace;
    vec3 lightPosition_cameraSpace = (V * M * vec4(lightPosition_worldSpace, 1)).xyz;
    lightDirection_cameraSpace = lightPosition_cameraSpace + eyeDirection_cameraSpace;

    vertexNormal_cameraSpace = (V * M * vec4(position_modelSpace, 0)).xyz;
}
