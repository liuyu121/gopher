package counter

type alertCounter int

// 工厂函数
func New(value int) alertCounter {
    return alertCounter(value)
}