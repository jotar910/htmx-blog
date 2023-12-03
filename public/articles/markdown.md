<div class="article__data">
<p><a href="https://www.dolthub.com/blog/2023-11-01-announcing-doltgresql/">DoltgreSQL</a> builds on top of the version-controlled database Dolt to provide Git-like log, diff, branch, and merge functionality for your Postgres database schema and data.</p>

<p>Dolt was born as an SQL database you can clone, fork, branch, and merge just like a Git repository. Using Dolt, application developers can build branch- and merge-workflows for their customers, for example, by sending pull requests to fix mistakes in the data. Similarly, Dolt enables a simple model to change a production database by branching it, applying the changes, then testing it in your staging setup, and eventually deploying back to production.</p>

<p>Since its beginnings, Dolt adopted MySQL's syntax and a <a href="https://docs.dolthub.com/cli-reference/git-comparison">command line-oriented paradigm</a> that is surely familiar to Git users.</p>

<p>DoltgreSQL focuses instead on the database server experience, providing a customizable, easy-to-deploy server. Furthermore, the company does not provide command-line support to better align with the general PostgreSQL audience:</p>

<blockquote>
<p>DoltgreSQL works by emulating a PostgreSQL server, and converting received commands into an AST that is given to an underlying Dolt server. This enabled us to get something up and running quickly, while leveraging the capabilities and functionality that Dolt already provides.</p>
</blockquote>

<p>This approach has the advantage of building these new features on top of Dolt's foundation, leveraging the latter's stability and reliability and reducing development scope and effort.</p>

<p>DoltHub says they investigated distinct approaches, including writing a foreign data wrapper, building a new PostgreSQL storage backend, and even forking PostgreSQL itself. Some of those approaches turned out to be too limited, while others, such as forking PostgreSQL, would have required years of development.</p>

<p>On the negative side, one shortcoming of the emulation-based approach is that you are not running the actual PostgreSQL binary. Instead, as mentioned, DoltgreSQL converts PostgreSQL syntax into its AST representation and runs it within the Dolt layer.</p>

<p>Once you have DoltgreSQL installed, you can connect to it using the <code>psql</code> command line client. To see the status of the database, you can run the query:</p>

<p><code>select * from dolt_status;</code></p>

<p>This will list all existing tables and specify whether they are new, staged, and so on. To add a table to the staging area, you run:</p>

<p><code>call dolt_add('my_table_name');</code></p>

<p>And commit changes with:</p>

<p><code>call dolt_commit('-m', 'updated schema');</code>.</p>

<p>The equivalent to <code>git log</code> is <code>select * from dolt_log;</code>.</p>

<p>Doltgres is still experimental and has several <a href="https://github.com/dolthub/doltgresql#limitations">limitations</a>, including lack of support on DoltHub and DoltLab, no authentication or user management, limited support for SSL connections, no support for replication, clustering, etc., and more.</p>

<p>While Dolt's "Git for data" value proposition may sound compelling, database expert <a href="http://www.jandrewrogers.com/about/">J. Andrew Rogers</a> observed on Hacker News this goal is <a href="https://news.ycombinator.com/item?id=31852067">not so unlike what multi-version concurrency control (MVCC) has&nbsp;attempted for decades</a> with several major drawbacks. Dolt CEO Tim Sehn highlighted <a href="https://docs.dolthub.com/sql-reference/benchmarks/latency">running Dolt against native MySQL shows it is only marginally slower for the <code>sysbench</code> benchmark</a>.</p>

<div class="author-section-full"> <!-- main wrapper for authors section -->
        <h2>About the Author</h2> <!-- section title -->
            <div class="author" data-id="author-Sergio-De-Simone"> <!-- main wrapper for each author -->
                <a href="/profile/Sergio-De-Simone/" class="avatar author__avatar"><img src="https://cdn.infoq.com/statics_s1_20231128091129/images/profiles/NovciOoQOAYWqYqRQBFo97SuMm0xbUiC.jpg"></a>
                <div class="content-author">
                    <h4><strong>Sergio De Simone</strong></h4>
                    <div class="show-author-bio">
                        <p>
                            <!-- author bio will be inserted by frontend -->
                        <b>Sergio De Simone</b> is a software engineer. Sergio has been working as a software engineer for over twenty five years across a range of different projects and companies, including such different work environments as Siemens, HP, and small startups. For the last 10+ years, his focus has been on development for mobile platforms and related technologies. He is currently working for BigML, Inc., where he leads iOS and macOS development.</p>
                        <span>
                            <div class="icon button-icon icon__plus-circle"></div><span class="show-more">Show more</span><span class="show-less">Show less</span>
                        </span>
                    </div>
                </div>
            </div>
    </div>
</div>
