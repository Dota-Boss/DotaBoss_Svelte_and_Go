<script>
  let teamA = '', teamB = '', result = null;

  async function scout() {
    const res = await fetch('http://localhost:8080/api/scout', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ teamA, teamB })
    });
    result = await res.json();
  }
</script>

<main>
  <h1>Dota Scout Reset 🔄</h1>
  <input bind:value={teamA} placeholder="Team A" />
  <input bind:value={teamB} placeholder="Team B" />
  <button on:click={scout}>Scout!</button>

  {#if result}
    <p>Matchup Rating: {result.matchupRating}</p>
    <p>Suggested Ban: {result.suggestedBan}</p>
  {/if}
</main>
