import pulumi
import pulumi_dagsterplus as dagsterplus

team = dagsterplus.Team("myTeam", name="pulumi-test-team")

pulumi.export("team_name", team.name)
