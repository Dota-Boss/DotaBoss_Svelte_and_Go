<script>
  let teamA = "";
  let teamB = "";
  let report = null;

  const scout = async () => {
    console.log("Sending:", { teamA, teamB });

    const res = await fetch("http://localhost:8080/api/scout", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ teamA, teamB }),
    });
    report = await res.json();
  };
</script>

<main>
  <h1>Dota Scout</h1>
  <form on:submit|preventDefault={scout}>
    <input bind:value={teamA} placeholder="Your Team" />
    <input bind:value={teamB} placeholder="Opponent Team" />
    <button type="submit">Scout!</button>
  </form>

  {#if report}
    <div class="report">
      <h2>Scouting Results</h2>
      <p><strong>Matchup Rating:</strong> {report.matchupRating}</p>
      <p><strong>Suggested Ban:</strong> {report.suggestedBan}</p>
    </div>
  {/if}
</main>
