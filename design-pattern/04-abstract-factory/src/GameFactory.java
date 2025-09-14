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
