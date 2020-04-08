#version 330

layout(location = 0) out vec4 colorOut;


// Vertex Normal (Sum of normals of all faces that use the vertex, in the special case of a cube this is equal to model coordinates)
// (from Vertex Shader)
in vec3 vertexNormal_cameraSpace;
// Vector from Vertex to Light (from Vertex Shader)
in vec3 lightDirection_cameraSpace;
//Vector from Vertex to Eye / Camera (from Vertex Shader)
in vec3 eyeDirection_cameraSpace;

void main() {
    vec3 materialDiffuseColor = vec3(0.75, 0.75, 0);
    float lightPower = 25;
    float ambientLight = 0.1;

    vec3 lightColor = vec3(1, 1, 1) * lightPower;
    float distance = length(lightDirection_cameraSpace);

    // normalized vector from fragment to light
    vec3 l = lightDirection_cameraSpace / distance;
    // normalized normal vector
    vec3 n = normalize(vertexNormal_cameraSpace);

    // normalized vector from fragment to eye
    vec3 e = normalize(eyeDirection_cameraSpace);
    // incident light reflection vector
    vec3 r = reflect(-l, n);

    // cosine of angle between eye and reflected light
    float cosAlpha = clamp(dot(e, r), 0, 1);
    // cosine of angle between light and normal
    float cosTheta = clamp(dot(n, l), 0, 1);

    float distanceSquare = (distance * distance);

    colorOut =
    vec4(materialDiffuseColor * lightColor * cosTheta / distanceSquare, 1) + // diffuse lighting
    vec4(materialDiffuseColor * lightColor * pow(cosAlpha, 7) / distanceSquare, 1) + // specular lighting
    vec4(ambientLight * materialDiffuseColor, 1); // fake ambient light
}
