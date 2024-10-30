window.onload = function () {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    url: "./swagger.json",
    dom_id: '#swagger-ui',
    deepLinking: true,
    showCommonExtensions: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [

    ],
    layout: "BaseLayout",
  });

  //</editor-fold>
};
