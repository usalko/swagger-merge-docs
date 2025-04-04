package docs

var IndexHtml = []byte(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta
      name="description"
      content="SwaggerUI"
    />
    <title>API Documentation (Swagger)</title>
    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.20.1/swagger-ui.css" />
  </head>
  <body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5.20.1/swagger-ui-bundle.js" crossorigin></script>
  <script src="https://unpkg.com/@mairu/swagger-ui-apikey-auth-form@1.2.1/dist/swagger-ui-apikey-auth-form.js" crossorigin></script>
  <script>
    window.onload = function() {
      window.ui = SwaggerUIBundle({
        url: "docs/swagger.yaml",
        dom_id: '#swagger-ui',
        presets: [
          SwaggerUIBundle.presets.apis,
        ],
        layout: "BaseLayout",
        requestInterceptor: function(request) {
          request.headers['X-CSRFToken'] = "{{ csrf_token }}"
          return request;
        },
        plugins: [
          SwaggerUIApiKeyAuthFormPlugin,
        ],
        configs: {
          apiKeyAuthFormPlugin: {
            forms: {
              'Custom authentication': {
                fields: {
                  email: {
                    type: 'text',
                    label: 'email',
                  },
                  password: {
                    type: 'password',
                    label: 'password',
                  },
                  // ...
                },
                authCallback(values, callback) {
                  // // do login stuff
                  //  ...
                  // // on error
                  // callback('error message');
                  // 
                  // // on success
                  // callback(null, 'the api key here');

                  // example using SwaggerUIs fetch api
                  ui.fn.fetch({
                    url: '/api/v1/auth/login/',
                    method: 'post',
                    headers: {
                      Accept: 'application/json',
                      'Content-Type': 'multipart/form-data',
                      'X-CSRFToken': '{{ csrf_token }}',
                    },
                    body: Object.entries(values).reduce((d,e) => (d.append(...e),d), new FormData())
                  }).then(function (response) {
                    const json = JSON.parse(response.data);
                    if (json.authorization_header) {
                      callback(null, json.authorization_header);
                    } else {
                      callback('Authentication error (incorrect username or password)?)');
                    }
                  }).catch(function (err) {
                    console.log(err, Object.entries(err));
                    callback('Authentication error (see browser console log)');
                  });
                },
              }
            },
            localStorage: {
              'Custom authentication': {}
            }
          }
        }
      });
    };
  </script>
  </body>
</html>`)
