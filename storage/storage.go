package storage

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sync"

	"server-go/models"
)

const dataDir = "./data"

var (
	mu      sync.RWMutex
	games   []models.Game
	users   []models.User
	records []models.GameRecord
)

func Init() error {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return err
	}

	games = defaultGames()

	if err := loadJSON("users.json", &users); err != nil {
		users = []models.User{}
	}
	if err := loadJSON("records.json", &records); err != nil {
		records = []models.GameRecord{}
	}
	log.Printf("[storage] 已加载 游戏 %d 个，用户 %d 个，战绩 %d 条", len(games), len(users), len(records))
	return nil
}

func defaultGames() []models.Game {
	return []models.Game{
		{
			ID:           "minesweeper",
			Name:         "扫雷",
			Description:  "经典扫雷：左键揭开、右键插旗。含简单、中等、复杂三档盘面与雷数。",
			CoverImage:   "/covers/minesweeper.svg",
			Tags:         []string{"益智", "经典"},
			Difficulties: []string{"简单", "中等", "复杂"},
		},
		{
			ID:           "snake",
			Name:         "贪吃蛇",
			Description:  "吃食物变长，别撞墙和自己。三档难度对应不同场地大小与速度。",
			CoverImage:   "/covers/snake.svg",
			Tags:         []string{"休闲", "经典"},
			Difficulties: []string{"简单", "中等", "复杂"},
		},
		{
			ID:           "24points",
			Name:         "24点",
			Description:  "用四张牌凑 24。三档难度对应不同时长挑战。",
			CoverImage:   "/covers/24points.svg",
			Tags:         []string{"益智", "数学"},
			Difficulties: []string{"简单", "中等", "复杂"},
		},
		{
			ID:           "2048",
			Name:         "2048",
			Description:  "滑动合并数字。简单 3×3、中等 4×4、复杂 5×5，目标分数随盘面调整。",
			CoverImage:   "/covers/2048.svg",
			Tags:         []string{"益智", "数学"},
			Difficulties: []string{"简单", "中等", "复杂"},
		},
		{
			ID:           "memory",
			Name:         "记忆翻牌",
			Description:  "翻开找相同一对。三档难度对应不同对数。",
			CoverImage:   "/covers/memory.svg",
			Tags:         []string{"记忆", "益智"},
			Difficulties: []string{"简单", "中等", "复杂"},
		},
		{
			ID:           "whack-a-mole",
			Name:         "打地鼠",
			Description:  "限时点击地鼠。三档难度对应洞数、时长与出现速度。",
			CoverImage:   "/covers/whack-a-mole.svg",
			Tags:         []string{"反应", "休闲"},
			Difficulties: []string{"简单", "中等", "复杂"},
		},
		{
			ID:           "slide-puzzle",
			Name:         "数字华容道",
			Description:  "滑动方块复原顺序。三档对应 3×3、4×4、5×5 盘面。",
			CoverImage:   "/covers/slide-puzzle.svg",
			Tags:         []string{"益智", "经典"},
			Difficulties: []string{"简单", "中等", "复杂"},
		},
		{
			ID:           "reaction-test",
			Name:         "反应测试",
			Description:  "变绿后尽快点击，多回合累计得分。三档对应回合数与惩罚力度。",
			CoverImage:   "/covers/reaction-test.svg",
			Tags:         []string{"反应", "休闲"},
			Difficulties: []string{"简单", "中等", "复杂"},
		},
		{
			ID:           "tic-tac-toe",
			Name:         "井字棋",
			Description:  "先手 X 对战电脑 O。简单随机、中等会堵、复杂为极小化极大最优走法。",
			CoverImage:   "/covers/tic-tac-toe.svg",
			Tags:         []string{"益智", "对战"},
			Difficulties: []string{"简单", "中等", "复杂"},
		},
		{
			ID:           "tetris",
			Name:         "俄罗斯方块",
			Description:  "下落方块,旋转拼消。←/→ 平移、↑ 旋转、↓ 加速、空格硬降。三档对应起始速度与加速节奏。",
			CoverImage:   "/covers/tetris.svg",
			Tags:         []string{"经典", "消除"},
			Difficulties: []string{"简单", "中等", "复杂"},
		},
		{
			ID:           "breakout",
			Name:         "打砖块",
			Description:  "鼠标控制挡板,物理反弹小球清空砖墙。三档对应砖墙行数、挡板大小与球速。",
			CoverImage:   "/covers/breakout.svg",
			Tags:         []string{"反应", "物理"},
			Difficulties: []string{"简单", "中等", "复杂"},
		},
		{
			ID:           "wordle",
			Name:         "猜词",
			Description:  "按字母位置反馈颜色线索。简单 4 字母、中等 5 字母、复杂 6 字母且次数更少。",
			CoverImage:   "/covers/wordle.svg",
			Tags:         []string{"文字", "推理"},
			Difficulties: []string{"简单", "中等", "复杂"},
		},
		{
			ID:           "gomoku",
			Name:         "五子棋",
			Description:  "15×15 棋盘,你执黑先手。简单随机、中等会守必杀、复杂启用棋形评估。",
			CoverImage:   "/covers/gomoku.svg",
			Tags:         []string{"对战", "策略"},
			Difficulties: []string{"简单", "中等", "复杂"},
		},
	}
}

func loadJSON(filename string, v interface{}) error {
	data, err := os.ReadFile(filepath.Join(dataDir, filename))
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func saveJSON(filename string, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dataDir, filename), data, 0644)
}

// Games
func GetGames() []models.Game {
	mu.RLock()
	defer mu.RUnlock()
	return games
}

func GetGame(id string) (models.Game, bool) {
	mu.RLock()
	defer mu.RUnlock()
	for _, g := range games {
		if g.ID == id {
			return g, true
		}
	}
	return models.Game{}, false
}

// Users
func GetUsers() []models.User {
	mu.RLock()
	defer mu.RUnlock()
	result := make([]models.User, len(users))
	copy(result, users)
	return result
}

func GetUser(id string) (models.User, bool) {
	mu.RLock()
	defer mu.RUnlock()
	for _, u := range users {
		if u.ID == id {
			return u, true
		}
	}
	return models.User{}, false
}

func CreateUser(u models.User) error {
	mu.Lock()
	defer mu.Unlock()
	users = append(users, u)
	return saveJSON("users.json", users)
}

func DeleteUser(id string) error {
	mu.Lock()
	defer mu.Unlock()
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return saveJSON("users.json", users)
		}
	}
	return nil
}

// Records
func GetRecordsByUser(userID string) []models.GameRecord {
	mu.RLock()
	defer mu.RUnlock()
	var result []models.GameRecord
	for _, r := range records {
		if r.UserID == userID {
			result = append(result, r)
		}
	}
	return result
}

func CreateRecord(r models.GameRecord) error {
	mu.Lock()
	defer mu.Unlock()
	records = append(records, r)
	return saveJSON("records.json", records)
}

func GetUserStats(userID string) models.UserStats {
	mu.RLock()
	defer mu.RUnlock()

	user, _ := getUserUnlocked(userID)
	stats := models.UserStats{
		UserID:    userID,
		UserName:  user.Name,
		GameStats: []models.GameStat{},
	}

	gameMap := map[string]*models.GameStat{}
	for _, r := range records {
		if r.UserID != userID {
			continue
		}
		stats.TotalGames++
		stats.TotalTime += r.Duration
		stats.TotalScore += r.Score

		gs, ok := gameMap[r.GameID]
		if !ok {
			name := r.GameID
			for _, g := range games {
				if g.ID == r.GameID {
					name = g.Name
					break
				}
			}
			gameMap[r.GameID] = &models.GameStat{GameID: r.GameID, GameName: name}
			gs = gameMap[r.GameID]
		}
		gs.PlayCount++
		gs.TotalTime += r.Duration
		if r.Score > gs.BestScore {
			gs.BestScore = r.Score
		}
	}

	for _, gs := range gameMap {
		stats.GameStats = append(stats.GameStats, *gs)
	}
	return stats
}

func getUserUnlocked(id string) (models.User, bool) {
	for _, u := range users {
		if u.ID == id {
			return u, true
		}
	}
	return models.User{}, false
}
