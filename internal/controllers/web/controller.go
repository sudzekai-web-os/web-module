package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sudzekai-web-os/core"
)

type Controller struct {
	configuration core.IConfiguration
	registry      core.IHandlersRegistry
}

func NewController(
	conf core.IConfiguration,
	registry core.IHandlersRegistry,
) *Controller {
	return &Controller{
		configuration: conf,
		registry:      registry,
	}
}

func (ctr *Controller) AddRoutes(registry core.IHandlersRegistry) {
	registry.AddNoFilterHandler("GET /web", ctr.GetPage)
}

func (ctr *Controller) GetPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	configurationBlocks := make([]string, 0)

	for key, value := range ctr.configuration.GetOptions() {
		configurationBlocks = append(
			configurationBlocks,
			fmt.Sprintf(`
			<div class="config-item">
				<span class="key">%s</span>
				<span class="value">%v</span>
			</div>
		`, key, value()),
		)
	}

	routeBlocks := make([]string, 0)

	for _, route := range ctr.registry.GetRoutes() {
		parts := strings.SplitN(route, " ", 2)

		if len(parts) != 2 {
			continue
		}

		methodClass := strings.ToLower(parts[0])

		routeBlocks = append(
			routeBlocks,
			fmt.Sprintf(`
				<div class="endpoint">
					<div class="endpoint-header">
						<span class="method %s">%s</span>
						<span class="path">%s</span>
					</div>

					<div class="description">
						API endpoint
					</div>
				</div>
			`, methodClass, parts[0], parts[1]),
		)
	}

	page := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <title>Sudzekai Web OS</title>

    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            min-height: 100vh;
            background: #121314;
            color: #f1f1f1;
            font-family: Arial, sans-serif;
        }

        .container {
            width: min(1100px, calc(100%% - 40px));
            margin: 0 auto;
            padding: 50px 0;
        }

        .header {
            margin-bottom: 40px;
        }

        .header h1 {
            font-size: 32px;
            font-weight: 600;
        }

        .header p {
            margin-top: 8px;
            color: #8e9297;
            font-size: 15px;
        }

        .grid {
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 20px;
        }

        .card {
            background: #1a1b1d;
            border: 1px solid #292b2e;
            border-radius: 12px;
            padding: 24px;
        }

        .card h2 {
            font-size: 18px;
            font-weight: 500;
            margin-bottom: 20px;
        }

        .config {
            display: flex;
            flex-direction: column;
            gap: 12px;
        }

        .config-item {
            display: flex;
            justify-content: space-between;
            align-items: center;

            padding: 12px 14px;
            background: #141516;
            border-radius: 8px;
        }

        .config-item .key {
            color: #8e9297;
            font-size: 14px;
        }

        .config-item .value {
            font-family: monospace;
            font-size: 14px;
            color: #d9dde1;
        }

        .endpoint {
            padding: 14px;
            background: #141516;
            border-radius: 8px;
            margin-bottom: 10px;
        }

        .endpoint:last-child {
            margin-bottom: 0;
        }

        .endpoint-header {
            display: flex;
            align-items: center;
            gap: 10px;
            margin-bottom: 7px;
        }

		.method {
			padding: 4px 7px;
			border-radius: 5px;
			font-family: monospace;
			font-size: 12px;
		}

		.method.get {
			background: #16351f;
			color: #6ee7a0;
		}

		.method.post {
			background: #352d16;
			color: #f5c96a;
		}

		.method.put {
			background: #162d3d;
			color: #70c8ff;
		}

		.method.patch {
			background: #30203d;
			color: #d39aff;
		}

		.method.delete {
			background: #3d1c1c;
			color: #ff7777;
		}

		.method.head {
			background: #25282b;
			color: #b9bec4;
		}

		.method.options {
			background: #25282b;
			color: #b9bec4;
		}

        .path {
            font-family: monospace;
            font-size: 14px;
        }

        .description {
            color: #777b80;
            font-size: 13px;
        }

        .status {
            display: inline-flex;
            align-items: center;
            gap: 7px;
            margin-top: 30px;
            color: #85898e;
            font-size: 13px;
        }

        .status-dot {
            width: 7px;
            height: 7px;
            border-radius: 50%%;
            background: #62c370;
        }

        @media (max-width: 700px) {
            .grid {
                grid-template-columns: 1fr;
            }

            .container {
                width: min(100%% - 24px, 1100px);
                padding: 30px 0;
            }
        }
    </style>
</head>

<body>

<div class="container">

    <header class="header">
        <h1>sudzekai web os</h1>
        <p>Web interface for Linux server administration</p>

        <div class="status">
            <span class="status-dot"></span>
            Server is running
        </div>
    </header>

    <main class="grid">

        <section class="card">
            <h2>Configuration</h2>

            <div class="config">
				%s
            </div>
        </section>


        <section class="card">
            <h2>Endpoints</h2>
				%s
        </section>

    </main>

</div>

</body>
</html>
`, strings.Join(configurationBlocks, "\n"), strings.Join(routeBlocks, "\n"))
	_, _ = w.Write([]byte(page))
}
