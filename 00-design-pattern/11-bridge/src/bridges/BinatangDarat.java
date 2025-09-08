package bridges;

import parent.Binatang;

public abstract class BinatangDarat implements Binatang {
    @Override
    public boolean hidupDiDarat() {
        return true;
    }

    @Override
    public boolean hidupDiLaut() {
        return false;
    }

    public abstract int getJumlahKaki();
}
