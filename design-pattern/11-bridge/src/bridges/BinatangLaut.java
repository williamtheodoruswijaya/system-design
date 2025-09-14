package bridges;

import parent.Binatang;

public abstract class BinatangLaut implements Binatang {
    @Override
    public boolean hidupDiLaut() {
        return true;
    }

    @Override
    public boolean hidupDiDarat() {
        return false;
    }
}
