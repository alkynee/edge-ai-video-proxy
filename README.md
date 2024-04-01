
# edge-ai-video-proxy\n\n![GitHub release (latest by date)](https://img.shields.io/github/v/release/chryscloud/video-edge-ai-proxy) \n\nEdge AI Video Proxy ingests multiple RTSP camera streams and provides a common interface for conducting AI operations on or near the Edge.\n\n## Why use Edge-AI Video Proxy?\nThe edge-ai-video-proxy is an easy to use collection mechanism from multiple cameras onto a single more powerful computer. A network of CCTV RTSP enabled cameras can be accessed through a simple GRPC interface, where Machine Learning algorithms can perform various Computer Vision tasks. Moreover, interesting footage can be annotated, selectively streamed and stored through a simple API for later analysis, computer vision tasks in the cloud or enriching the Machine Learning training samples.\n\n<p align="center">\n<img src="https://storage.googleapis.com/chrysaliswebassets/chrysalis-video-edge-ai-proxy.png" title="Chrysalis Cloud" />\n<p align="center">\n\n## Documentation\nYou can find more extensive documentation [here](https://chryscloud.github.io/api_doc/)\n\n## Contents\n* [Prerequisites](#prerequisites)\n* [Quick Start](#quick-start)\n* [Usage](#usage)\n* [Examples](#examples)\n* [Custom configuration](#custom-configuration)\n* [Building from source](#building-from-source)\n\n## Prerequisites\n- [Docker](https://docs.docker.com/engine/install/)\n- [Docker Compose](https://docs.docker.com/compose/install/)\nRead specific configuration options [here](https://chryscloud.github.io/api_doc/edge-proxy/getting-started/prerequisites/)\n\n## Quick Start\nBy default edge-ai-video-proxy requires these ports:\n- *8905* for web portal\n- *8909* for RESTful API (portal API)\n- *50001* for client grpc connection\n- *6379* for redis\nMake sure before your run it that these ports are available.\n```\ncurl -O https://raw.githubusercontent.com/chryscloud/api_doc/master/install-chrysedge.sh\n\n# Give exec permission\nchmod 777 install-chrysedge.sh\n\n# run installation script\n./install-chrysedge.sh\n```\nStart the docker images:\n```python\ndocker-compose up\n\n# or to run it in daemon mode:\n\ndocker-compose up -d\n```\nOpen browser and visit `chrysalisportal` at address: `http://localhost:8905`\nFor installation outside of WSL 2 on Windows please check manual installation steps [here](https://chryscloud.github.io/api_doc/edge-proxy/getting-started/quick-start/#manual-installation)\n\n## Usage\nOpen your browser and go to: `http://localhost:8905`\nOn the first visit Edge Proxy will display a RTSP docker container icon. Click on it. This will initiate the pull for the latest version of the docker container pre-compiled to be used with RTSP enabled cameras.\n\n## Authors\n- **alkynee** - Current work - [AlkyneTech](https://alkynetech.com)\n\n# Contributors\nFind out how to contribute by reading `CONTRIBUTING.md`.\n\n# Versioning\nCurrent version is initial release - v0.0.8 prerelease\n\n# License\nThis project is licensed under the Apache 2.0 License. See `LICENSE` for details.