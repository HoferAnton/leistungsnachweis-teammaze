#version 330
//Model View Projection Matrix. Projects Model Space onto Screen Space
uniform mat4 MVP;
//Model Matrix. Projects Model Space onto World Space
uniform mat4 M;
//View Matrix. Projects World Space into Camera Space
uniform mat4 V;

//Light Position in World Space
uniform vec3 lightPosition_worldSpace;

in vec3 position_modelSpace;

// Vertex Normal (Sum of normals of all faces that use the vertex, in the special case of a cube this is equal to model coordinates)
out vec3 vertexNormal_cameraSpace;
// Vector from Vertex to Light
out vec3 lightDirection_cameraSpace;
//Vector from Vertex to Eye / Camera
out vec3 eyeDirection_cameraSpace;

void main() {
    gl_Position = MVP * vec4(position_modelSpace, 1);

    vec3 position_worldSpace = (M * vec4(position_modelSpace, 1)).xyz;
    vec3 position_cameraSpace = (V * M * vec4(position_modelSpace, 1)).xyz;

    eyeDirection_cameraSpace = vec3(0, 0, 0) - position_cameraSpace;
    vec3 lightPosition_cameraSpace = (V *vec4(lightPosition_worldSpace, 1)).xyz;
    lightDirection_cameraSpace = lightPosition_cameraSpace + eyeDirection_cameraSpace;

    vertexNormal_cameraSpace = (V * M * vec4(position_modelSpace, 0)).xyz;
}
