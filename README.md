# Buzzer Blog

Buzzer is a Golang-based project featuring a blog and a CMS backoffice, highlighting HTMX for responsiveness and speed.
It emphasizes the decoupling of content from technology, using Markdown and HTML custom elements, making it versatile for integration with various web technologies or frameworks.

## Features
- **HTMX Integration:** For dynamic, fast-loading web pages.
- **Markdown for Posts:** Ensures ease of writing and formatting.
- **Custom HTML Elements:** Enhance the interactivity and user experience.
- **Tailwind CSS:** Utilizes Tailwind for responsive and modern UI design.
- **SQLite Database:** Lightweight and easy to manage.

## Setup
Clone the repository:
```bash
git clone https://github.com/jotar910/htmx-blog.git
```

## Running
To run the Buzzer application in both the `/blog` and `/cms` directories, follow these steps:
- Start the Go server using either `air` or `go run .`.
- To compile TypeScript, Tailwind, and Go Templ, execute `npm run watch`.

## Building
To build the Buzzer project, in each of the `/blog` and `/cms` directories:
- Run `go build` to compile the Go code.
- Execute `npm run build` to compile TypeScript, Tailwind, and Go Templ.
- The output will be in the `dist/` directory. Copy the `public` folder to this directory to serve the built application.

## License
This project is licensed under the [GNU General Public License (GPL)](https://www.gnu.org/licenses/gpl-3.0.en.html). The GPL is a copyleft license that ensures any derivative work is also open source under the same terms.

