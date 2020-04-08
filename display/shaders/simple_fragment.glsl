#version 330

layout(location = 0) out vec4 colorOut;


in vec3 vertexNormal_cameraSpace;
in vec3 lightDirection_cameraSpace;
in vec3 eyeDirection_cameraSpace;

void main() {
    vec3 materialDiffuseColor = vec3(0.75, 0.75, 0);
    float lightPower = 25;
    float ambientLight = 0.1;

    vec3 lightColor = vec3(1, 1, 1) * lightPower;
    float distance = length(lightDirection_cameraSpace);

    vec3 l = lightDirection_cameraSpace / distance;
    vec3 n = normalize(vertexNormal_cameraSpace);

    vec3 e = normalize(eyeDirection_cameraSpace);
    vec3 r = reflect(-l, n);

    float cosAlpha = clamp(dot(e, r), 0, 1);
    float cosTheta = clamp(dot(n, l), 0, 1);

    float distanceSquare = (distance * distance);

    colorOut =
    vec4(materialDiffuseColor * lightColor * cosTheta / distanceSquare, 1) +
    vec4(materialDiffuseColor * lightColor * pow(cosAlpha, 7) / distanceSquare, 1) +
    vec4(ambientLight * materialDiffuseColor, 1);
}
