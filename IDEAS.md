# Little log of ideas to capture while they come up

5. look at wiki-go it might be the ideal md front end

6. cli command to move canvases between artists, e.g there could be a cloud functions artist that handles production, and we might want to move from a dev artist to production 

8. Maybe a way for allowing user to inherit default context changes? So as we refine the handling of AI agents, thne when we have an update cli, user can --pull-context, and a clever system tries to update their AGENTS.md files with the new ways to handle the AI's

9. It is complicated to setup the git repos at first, the AI i used got lost in the strcucture and what is required, maybe it needs special context notes, maybe even its own workflow file ?

10. integrate the git engine work as cli commands

COPY across the new git strategy make commands etc
COPY in readme snippets and AGENTS workflow snippets?

11. becasue the first task for a user is to setup their git config and or remotes, the subsequent cli commands could inherit the config from those and emulate it, so artist create could theoretically make a remote repo and push to it immediately, same for canvases, and for delete it could also delete the remote repos

12. need to stop ai's overwritn gthe established readme, gemini and makefile layouts

13. integrate memory bank https://kilocode.ai/docs/advanced-usage/memory-bank

14. rename all gemini files to AGENTS.md

15. look into taskmaster https://github.com/eyaltoledano/claude-task-master

16. on init check for gh cli in ssh, task-master cli if decided to use and other requirements

17. need to stop ai from ever querying klocal env variables as it reveals secrets in chat e.g cd artist-golang/canvas-salesIQ-api-wrapper && env | grep ZOHO

18. CLI first is a great parttern for rapid prototyping go, maybe the arist-golang can pre set this up with cobra, and then the makefile can stay as more generic build and tests, and cli driven commands for the bus logic

19. delete artist or canvas sometimes needs sudo for the rm -rf

