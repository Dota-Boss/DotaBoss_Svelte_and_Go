<script>
    let teamA = "";
    let teamB = "";
    let report = null;

    const scout = async () => {
        const res = await fetch("http://localhost:8080/api/scout", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ teamA, teamB })
        });
        report = await res.json();
    };
</script>

<main>
    <h1>Dota Scout</h1>
    <input bind:value={teamA} placeholder="Your Team" />
    <input bind:value={teamB} placeholder="Opponent Team" />
    <button on:click={scout}>Scout!</button>

    {#if report}
        <div class="report">
            <h2>Scouting Results</h2>
            <p><strong>Matchup Rating:</strong> {report.matchupRating}</p>
            <p><strong>Suggested Ban:</strong> {report.suggestedBan}</p>
        </div>
    {/if}
</main>
