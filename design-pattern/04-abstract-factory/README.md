# Abstract Factory

Ini adalah implementasi advanced dari Factory Design Pattern. Dimana dalam kasus ini, alih-alih menggunakan Factory Design Pattern biasa, ini kita gunakan ketika ada kasus dimana ada 2 class yang dependent satu sama lain.
Contoh:
```java
public class Game {
    public Level level;
    public Arena arena;
    
    public Game(Level level, Arena arena) {
        this.level = level;
        this.arena = arena;
    }

    public void start() {
        level.start();
        arena.start();
    }
}
```
Dimana level dan arena adalah sebuah interface seperti ini:

- Level:
```java
interface Level {
    public void start();
}

class LevelEasy implements Level {
    @Override
    public void start() {
        System.out.println("Level Easy Start");
    }
}

class LevelMedium implements Level {
    @Override
    public void start() {
        System.out.println("Level Medium Start");
    }
}

class LevelHard implements Level {
    @Override
    public void start() {
        System.out.println("Level Hard Start");
    }
}
```

- Arena:
```java
interface Arena {
    public void start();
}

class ArenaEasy implements Arena{
    @Override
    public void start() {
        System.out.println("Arena Easy");
    }
}

class ArenaMedium implements Arena{
    @Override
    public void start() {
        System.out.println("Arena Medium");
    }
}

class ArenaHard implements Arena{
    @Override
    public void start() {
        System.out.println("Arena Hard");
    }
}
```
Nah, di Main, kurang lebih kan bakal kek gini yah:
```java
public class Main {
    public static void main(String[] args) {
        Game gameEasy = new Game(new LevelEasy(), new ArenaEasy());
        gameEasy.start();

        Game gameMedium = new Game(new LevelMedium(), new ArenaMedium());
        gameMedium.start();

        Game gameHard = new Game(new LevelHard(), new ArenaMedium());
        gameHard.start();
    }
}
```
Dimana ini akan create sebuah dependency antara Level dan Arena. Dimana pada game easy, harus Level Easy dan Arena Easy yang kita initiate, tapi kalau initiate manual seperti diatas, dimana Level Easy dan Arena Medium bisa aja ke initiate. Nah ini jelas jadi masalah, oleh karena itu, kita bisa pakai Abstract Factory dimana kita wrap 2 class sekaligus untuk memastikan dependency yang sesuai.
Contoh class GameFactory:
```java
interface GameFactory {
    public Level createLevel();
    public Arena createArena();
}

class GameFactoryEasy implements GameFactory {
    @Override
    public Level createLevel() {
        return new LevelEasy();
    }

    @Override
    public Arena createArena() {
        return new ArenaEasy();
    }
}

class GameFactoryMedium implements GameFactory {
    @Override
    public Level createLevel() {
        return new LevelMedium();
    }

    @Override
    public Arena createArena() {
        return new ArenaMedium();
    }
}

class GameFactoryHard implements GameFactory {
    @Override
    public Level createLevel() {
        return new LevelHard();
    }

    @Override
    public Arena createArena() {
        return new ArenaHard();
    }
}
```
Nah, terus untuk Constructor Gamenya, sekarang alih-alih pakai 2 parameter, kita pakai aja langsung si Factory ini.
```java
public class Game {
    public Level level;
    public Arena arena;

//    Cara yang salah
    public Game(Level level, Arena arena) {
        this.level = level;
        this.arena = arena;
    }

//    Cara yang benar (setiap perubahan tinggal tambahin dibawah tanpa perlu adjust parameter constructor lagi)
    public Game(GameFactory game) {
        this.level = game.createLevel();
        this.arena = game.createArena();
    }
}
```
Nanti di main kita bisa langsung panggil kek gini:
```java
public class Main {
    public static void main(String[] args) {
//        Cara yang salah
//        Game gameEasy = new Game(new LevelEasy(), new ArenaEasy());
//        gameEasy.start();
//
//        Game gameMedium = new Game(new LevelMedium(), new ArenaMedium());
//        gameMedium.start();
//
//        Game gameHard = new Game(new LevelHard(), new ArenaMedium());
//        gameHard.start();

        Game gameEasy = new Game(new GameFactoryEasy());
        gameEasy.start();

        Game gameMedium = new Game(new GameFactoryMedium());
        gameMedium.start();

        Game gameHard = new Game(new GameFactoryHard());
        gameHard.start();
    }
}
```
Dengan begini, maka tidak akan ada kekeliruan dimana Game level easy tapi ada arena hard. Ini juga mempermudah penambahan atribut seperti misal, kita mau nambahin Enemy.
```java
interface Enemy {
    public void start();
}

class EnemyEasy implements Enemy {
    @Override
    public void start() {
        System.out.println("Enemy Easy");
    }
}

class EnemyMedium implements Enemy {
    @Override
    public void start() {
        System.out.println("Enemy Medium");
    }
}

class EnemyHard implements Enemy {
    @Override
    public void start() {
        System.out.println("Enemy Hard");
    }
}
```
Nah, kita tidak perlu update parameter, tapi langsung aja di factorynya:
```java
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

interface GameFactory {
    public Level createLevel();
    public Arena createArena();

    // tambahan Enemy (kalau ada tambahan modif attribute, tinggal tambahin disini)
    public Enemy createEnemy();
}

class GameFactoryEasy implements GameFactory {
    @Override
    public Level createLevel() {
        return new LevelEasy();
    }

    @Override
    public Arena createArena() {
        return new ArenaEasy();
    }

    @Override
    public Enemy createEnemy() {
        return new EnemyEasy();
    }
}

class GameFactoryMedium implements GameFactory {
    @Override
    public Level createLevel() {
        return new LevelMedium();
    }

    @Override
    public Arena createArena() {
        return new ArenaMedium();
    }

    @Override
    public Enemy createEnemy() {
        return new EnemyMedium();
    }
}

class GameFactoryHard implements GameFactory {
    @Override
    public Level createLevel() {
        return new LevelHard();
    }

    @Override
    public Arena createArena() {
        return new ArenaHard();
    }

    @Override
    public Enemy createEnemy() {
        return new EnemyHard();
    }
}
```