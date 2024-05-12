<p align="center"><a href="https://www.codechefvit.com" target="_blank"><img src="https://i.ibb.co/4J9LXxS/cclogo.png" width=160 title="CodeChef-VIT" alt="Codechef-VIT"></a>
</p>
<br />

# Devsoc Backend '24

The official Backend API for DEVSOC'24 Hackathon Portal

## Tech Stack

- [Golang](https://go.dev): Docker is a platform for developing, shipping, and running applications using containers.

  - [Echo](https://echo.labstack.com/): A fast http router based on the go standard HTTP.
  - [pgx](https://github.com/jackc/pgx): A postgresql driver for go.
  - [gomail](https://github.com/go-gomail/gomail): A library to send emails using go.
  - [go-redis](https://github.com/redis/go-redis): Go redis client.

- [Docker](https://www.docker.com): Docker is a platform for developing, shipping, and running applications using containers.
- [PostgreSQL](https://www.postgresql.org/): The worlds most advanced Open Source Relational Database.
- [Redis](https://redis.io): The world’s fastest in-memory database from the ones who built it.

## Features

- User Authentication
- Teams
- Project management
- Admin routes where Admins can directly score teams

# How To Run

## Prerequisites:

- [Docker](https://www.docker.com): .
- [Postman](https://www.postman.com)/[Apidog](https://apidog.com): A tool to test backend APIs without having to write frontends.
- [goose](https://github.com/pressly/goose): Goose is a database migration tool. Manage your database schema by creating incremental SQL changes or Go functions.

## Steps

1.  Clone the Repository: `git clone https://github.com/CodeChefVIT/devsoc-backend-24`
2.  Start the containers: `cd devsoc-backend-24 && cp .env.example .env && make build`. Please ensure that you put the correct SMTP credentials to get email services
3.  Run the migrations: `make migrate-up`
4.  Use postman and test the api at the endpoint `http://localhost/api`

# Authors

<table>
<tr align="center">
<td>
<p align="center">
<img src="https://avatars.githubusercontent.com/u/71623796?v=4" width="200" height="200" alt="Vedant"
style="border: 2px solid grey; width: 170px; height: 170px" />
</p>
<p style="font-size: 17px; font-weight: 600">Vedant Matanhelia</p>
<p align="center">
<a href="https://github.com/DanglingDynamo"><img
src="http://www.iconninja.com/files/241/825/211/round-collaboration-social-github-code-circle-network-icon.svg"
width="36" height="36" alt="GitHub" /></a>
<a href="https://www.linkedin.com/in/vedant-matanhelia-aa171027b/">
<img src="http://www.iconninja.com/files/863/607/751/network-linkedin-social-connection-circular-circle-media-icon.svg"
width="36" height="36" alt="LinkedIn" />
</a>
</p>
</td>

<td>
<p align="center">
<img src="https://avatars.githubusercontent.com/u/86644389?v=4" width="200" height="200" alt="Aman"
style="border: 2px solid grey; width: 170px; height: 170px" />
</p>
<p style="font-size: 17px; font-weight: 600">Aman L</p>
<p align="center">
<a href="https://github.com/Killerrekt"><img
src="http://www.iconninja.com/files/241/825/211/round-collaboration-social-github-code-circle-network-icon.svg"
width="36" height="36" alt="GitHub" /></a>
<a href="https://www.linkedin.com/in/aman-l-922819251/">
<img src="http://www.iconninja.com/files/863/607/751/network-linkedin-social-connection-circular-circle-media-icon.svg"
width="36" height="36" alt="LinkedIn" />
</a>
</p>
</td>

<td>
<p align="center">
<img src="https://avatars.githubusercontent.com/u/133687995?v=4" width="200" height="200" alt="Prateek"
style="border: 2px solid grey; width: 170px; height: 170px" />
</p>
<p style="font-size: 17px; font-weight: 600">Prateek Srivastava</p>
<p align="center">
<a href="https://github.com/prateek-srivastava001"><img
src="http://www.iconninja.com/files/241/825/211/round-collaboration-social-github-code-circle-network-icon.svg"
width="36" height="36" alt="GitHub" /></a>
<a href="https://www.linkedin.com/in/prateeksrivastava-/">
<img src="http://www.iconninja.com/files/863/607/751/network-linkedin-social-connection-circular-circle-media-icon.svg"
width="36" height="36" alt="LinkedIn" />
</a>
</p>
</td>

<td>
<p align="center">
<img src="https://avatars.githubusercontent.com/u/84951451?v=4" width="200" height="200" alt="Akshat"
style="border: 2px solid grey; width: 170px; height: 170px" />
</p>
<p style="font-size: 17px; font-weight: 600">Akshat Gupta</p>
<p align="center">
<a href="https://github.com/Oik17"><img
src="http://www.iconninja.com/files/241/825/211/round-collaboration-social-github-code-circle-network-icon.svg"
width="36" height="36" alt="GitHub" /></a>
<a href="https://www.linkedin.com/in/akshat-gupta-864b39235/">
<img src="http://www.iconninja.com/files/863/607/751/network-linkedin-social-connection-circular-circle-media-icon.svg"
width="36" height="36" alt="LinkedIn" />
</a>
</p>
</td>

</tr>

<tr align="center">
<td>
<p align="center">
<img src="https://avatars.githubusercontent.com/u/50650788?v=4" width="200" height="200" alt="Shivam"
style="border: 2px solid grey; width: 170px; height: 170px" />
</p>
<p style="font-size: 17px; font-weight: 600">Shivam Sharma</p>
<p align="center">
<a href="https://github.com/Mr-Emerald-Wolf"><img
src="http://www.iconninja.com/files/241/825/211/round-collaboration-social-github-code-circle-network-icon.svg"
width="36" height="36" alt="GitHub" /></a>
<a href="https://www.linkedin.com/in/shivam-sharma-6a0b1b1a7/">
<img src="http://www.iconninja.com/files/863/607/751/network-linkedin-social-connection-circular-circle-media-icon.svg"
width="36" height="36" alt="LinkedIn" />
</a>
</p>
</td>

<td>
<p align="center">
<img src="https://avatars.githubusercontent.com/u/91564450?v=4" width="200" height="200" alt="Aaditya"
style="border: 2px solid grey; width: 170px; height: 170px" />
</p>
<p style="font-size: 17px; font-weight: 600">Aaditya Mahanta</p>
<p align="center">
<a href="https://github.com/aditansh"><img
src="http://www.iconninja.com/files/241/825/211/round-collaboration-social-github-code-circle-network-icon.svg"
width="36" height="36" alt="GitHub" /></a>
<a href="https://www.linkedin.com/in/aadityamahanta/">
<img src="http://www.iconninja.com/files/863/607/751/network-linkedin-social-connection-circular-circle-media-icon.svg"
width="36" height="36" alt="LinkedIn" />
</a>
</p>
</td>

<td>
<p align="center">
<img src="https://avatars.githubusercontent.com/u/100862487?v=4" width="200" height="200" alt="Shivam"
style="border: 2px solid grey; width: 170px; height: 170px" />
</p>
<p style="font-size: 17px; font-weight: 600">Shivam Gutgutia</p>
<p align="center">
<a href="https://github.com/shivamgutgutia"><img
src="http://www.iconninja.com/files/241/825/211/round-collaboration-social-github-code-circle-network-icon.svg"
width="36" height="36" alt="GitHub" /></a>
<a href="https://www.linkedin.com/in/shivamgutgutia/">
<img src="http://www.iconninja.com/files/863/607/751/network-linkedin-social-connection-circular-circle-media-icon.svg"
width="36" height="36" alt="LinkedIn" />
</a>
</p>
</td>
</tr>
</table>

# License

Copyright © 2024, [CodeChef-VIT](https://github.com/CodeChefVIT) and all other contributors.
Released under the [MIT License](LICENSE).

<p align="center">
Made with :heart: by <a href="https://www.codechefvit.com" target="_blank">CodeChef-VIT</a>
</p>
