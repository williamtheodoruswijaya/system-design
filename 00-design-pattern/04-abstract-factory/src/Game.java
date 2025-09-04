public class Game {
    public Level level;
    public Arena arena;
    public Enemy enemy;

//    Cara yang salah
    public Game(Level level, Arena arena) { // <- enemy harus di update disini kalau gapake GameFactory
        this.level = level;
        this.arena = arena;
    }

//    Cara yang benar (setiap perubahan tinggal tambahin dibawah tanpa perlu adjust parameter constructor lagi)
    public Game(GameFactory game) {
        this.level = game.createLevel();
        this.arena = game.createArena();
        this.enemy = game.createEnemy();
    }

    public void start() {
        level.start();
        arena.start();
        enemy.start();
    }
}
