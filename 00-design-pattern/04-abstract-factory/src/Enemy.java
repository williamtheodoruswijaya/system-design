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
