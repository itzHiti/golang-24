package internal

// Team represents an esports team
type Team struct {
	Name        string
	Country     string
	FoundedYear int
	Players     []Player
}

// Player represents a player in an esports team
type Player struct {
	Name  string
	Age   int
	Role  string
	Stats string
}

// GetTeams returns a list of esports teams
func GetTeams() []Team {
	// Replace with your actual esports teams data
	teams := []Team{
		{
			Name:        "Astralis",
			Country:     "Denmark",
			FoundedYear: 2016,
			Players: []Player{
				{"device", 25, "AWPer", "1.20 rating"},
				{"dupreeh", 28, "Entry Fragger", "1.15 rating"},
				// Add more players as needed
			},
		},
		// Add more teams as needed
	}
	return teams
}

// GetPlayerByName returns the profile of a specific player by name
func GetPlayerByName(teamName string, playerName string) (*Player, bool) {
	// Replace with your logic to find the player in the specified team
	teams := GetTeams()
	for _, team := range teams {
		if team.Name == teamName {
			for _, player := range team.Players {
				if player.Name == playerName {
					return &player, true
				}
			}
		}
	}
	return nil, false
}
