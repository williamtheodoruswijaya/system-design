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
