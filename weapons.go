package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/widget"
)

type weaponsDetail struct {
	missile     int
	beam        int
	pulse       int
	plasma      int
	fusion      int
	sandcaster  int
	accelerator int
}

var (
	missiles      = widget.NewLabel("Missile Turrets")
	beamLasers    = widget.NewLabel("Beam Laser Turrets")
	pulseLasers   = widget.NewLabel("Pulse Laser Turrets")
	fusion        = widget.NewLabel("Fusion Guns")
	sand          = widget.NewLabel("Sandcaster Turrets")
	plasma        = widget.NewLabel("Plasma Guns")
	particle      = widget.NewLabel("Particle Beams")
	weaponDetails = widget.NewVBox(missiles, beamLasers, pulseLasers, fusion, sand, particle)

	missilesSelect = widget.NewSelect(weaponLevel, nothing)
	beamSelect     = widget.NewSelect(weaponLevel, nothing)
	pulseSelect    = widget.NewSelect(weaponLevel, nothing)
	plasmaSelect   = widget.NewSelect(weaponLevel, nothing)
	fusionSelect   = widget.NewSelect(weaponLevel, nothing)
	sandSelect     = widget.NewSelect(weaponLevel, nothing)
	particleSelect = widget.NewSelect(weaponLevel, nothing)

	missileItem  = widget.NewFormItem("Missiles", missilesSelect)
	beamItem     = widget.NewFormItem("Beam Laser Turrets", beamSelect)
	pulseItem    = widget.NewFormItem("Pulse Laser Turrets", pulseSelect)
	fusionItem   = widget.NewFormItem("Fusion Guns", fusionSelect)
	sandItem     = widget.NewFormItem("Sand Caster Turrets", sandSelect)
	plasmaItem   = widget.NewFormItem("Plasma Guns", plasmaSelect)
	particleItem = widget.NewFormItem("Particle Beam Accelerators", particleSelect)

	weaponForm = widget.NewForm(
		missileItem, beamItem, pulseItem, fusionItem, sandItem, plasmaItem, particleItem,
	)

	weapons weaponsDetail = weaponsDetail{
		missile: 0, beam: 0, pulse: 0, plasma: 0, fusion: 0,
		sandcaster: 0, accelerator: 0,
	}
)

func (w *weaponsDetail) init(form *widget.Form, box *widget.Box) {
	missilesSelect.SetSelected("0")
	missilesSelect.OnChanged = w.missileChanged

	beamSelect.SetSelected("0")
	beamSelect.OnChanged = w.beamChanged

	pulseSelect.SetSelected("0")
	pulseSelect.OnChanged = w.pulseChanged

	fusionSelect.SetSelected("0")
	fusionSelect.OnChanged = w.fusionChanged

	plasmaSelect.SetSelected("0")
	plasmaSelect.OnChanged = w.plasmaChanged

	sandSelect.SetSelected("0")
	sandSelect.OnChanged = w.sandChanged

	particleSelect.SetSelected("0")
	particleSelect.OnChanged = w.particleChanged

	box.Children = append(box.Children, weaponDetails)

	sand.Hide()
	missiles.Hide()
	beamLasers.Hide()
	pulseLasers.Hide()
	fusion.Hide()
	sand.Hide()
	plasma.Hide()
	particle.Hide()

	form.AppendItem(widget.NewFormItem("Weapons", weaponForm))
}

func (w *weaponsDetail) missileChanged(value string) {
	missileLaunchers, err := strconv.Atoi(value)
	if err == nil {
		weapons.missile = missileLaunchers
		if w.countWeapons() > hull.maxHP {
			weapons.missile = missileLaunchers - w.countWeapons() + hull.maxHP
		} else if weapons.missile < 0 {
			weapons.missile = 0
		}
		if weapons.missile != missileLaunchers {
			missilesSelect.SetSelected(strconv.Itoa(weapons.missile))
		}
	}
	w.buildMissile()
	berths.buildCrew()
}

func (w *weaponsDetail) beamChanged(value string) {
	beamTurrets, err := strconv.Atoi(value)
	if err == nil {
		weapons.beam = beamTurrets
		if w.countWeapons() > hull.maxHP {
			weapons.beam = beamTurrets - w.countWeapons() + hull.maxHP
			if weapons.beam < 0 {
				weapons.beam = 0
			}
			beamSelect.SetSelected(strconv.Itoa(weapons.beam))
		}
	}
	w.buildBeam()
	if weapons.missile > 0 {
		missiles.SetText(fmt.Sprintf("Missile Turrets: %d", weapons.missile))
		missiles.Show()
	} else {
		missiles.Hide()
	}
	berths.buildCrew()
}

func (w *weaponsDetail) pulseChanged(value string) {
	pulse, err := strconv.Atoi(value)
	if err == nil {
		weapons.pulse = pulse
		if w.countWeapons() > hull.maxHP {
			weapons.pulse = pulse - w.countWeapons() + hull.maxHP
			if weapons.pulse < 0 {
				weapons.pulse = 0
			}
			pulseSelect.SetSelected(strconv.Itoa(weapons.pulse))
		}
	}
	w.buildPulse()
	berths.buildCrew()
}

func (w *weaponsDetail) fusionChanged(value string) {
	fusion, err := strconv.Atoi(value)
	if err == nil {
		weapons.fusion = fusion
		if w.countWeapons() > hull.maxHP {
			weapons.fusion = fusion - w.countWeapons() + hull.maxHP
			if weapons.fusion < 0 {
				weapons.fusion = 0
			}
			fusionSelect.SetSelected(strconv.Itoa(weapons.fusion))
		}
	}
	w.buildFusion()
	berths.buildCrew()
}

func (w *weaponsDetail) sandChanged(value string) {
	sand, err := strconv.Atoi(value)
	if err == nil {
		weapons.sandcaster = sand
		if w.countWeapons() > hull.maxHP {
			weapons.sandcaster = sand - w.countWeapons() + hull.maxHP
			if weapons.sandcaster < 0 {
				weapons.sandcaster = 0
			}
			sandSelect.SetSelected(strconv.Itoa(weapons.sandcaster))
		}
	}
	w.buildSand()
	berths.buildCrew()
}

func (w *weaponsDetail) plasmaChanged(value string) {
	plasma, err := strconv.Atoi(value)
	if err == nil {
		weapons.plasma = plasma
		if w.countWeapons() > hull.maxHP {
			weapons.plasma = plasma - w.countWeapons() + hull.maxHP
			if weapons.plasma < 0 {
				weapons.plasma = 0
			}
			plasmaSelect.SetSelected(strconv.Itoa(weapons.plasma))
		}
	}
	w.buildPlasma()
	berths.buildCrew()
}

func (w *weaponsDetail) particleChanged(value string) {
	particle, err := strconv.Atoi(value)
	if err == nil {
		weapons.accelerator = particle
		if w.countWeapons() > hull.maxHP {
			weapons.accelerator = particle - w.countWeapons() + hull.maxHP
			if weapons.accelerator < 0 {
				weapons.accelerator = 0
			}
			particleSelect.SetSelected(strconv.Itoa(weapons.accelerator))
		}
	}
	w.buildParticle()
	berths.buildCrew()
}

func (w *weaponsDetail) buildMissile() {
	if weapons.missile > 0 {
		missiles.SetText(
			buildAmmoWeaponString("Triple Missile turrets: %d, tons: %d, ammo tons: %d",
				weapons.missile, int(float32(weapons.missile)+0.9999), int(float32(4*weapons.missile)+0.9999)))
		missiles.Show()
	} else {
		missiles.Hide()
	}
	missiles.Refresh()
}

func (w *weaponsDetail) buildBeam() {
	if weapons.beam > 0 {
		beamLasers.SetText(
			buildWeaponString("Triple Beam laser turrets: %d, tons: %d",
				weapons.beam, int(float32(weapons.beam)+.9999)))
		beamLasers.Show()
	} else {
		beamLasers.Hide()
	}
	beamLasers.Refresh()
}

func (w *weaponsDetail) buildPulse() {
	if weapons.pulse > 0 {
		pulseLasers.SetText(
			buildWeaponString("Triple Pulse lasr turrets: %d, tons: %d",
				weapons.pulse, int(float32(weapons.pulse)+.9999)))
		pulseLasers.Show()
	} else {
		pulseLasers.Hide()
	}
	pulseLasers.Refresh()
}

func (w *weaponsDetail) buildPlasma() {
	if weapons.plasma > 0 {
		plasma.SetText(
			buildWeaponString("Double Plasma gun turrets: %d, tons: %d",
				weapons.plasma, int(float32(2*weapons.plasma)+.9999)))
		plasma.Show()
	} else {
		plasma.Hide()
	}
	plasma.Refresh()
}

func (w *weaponsDetail) buildFusion() {
	if weapons.fusion > 0 {
		fusion.SetText(
			buildWeaponString("Double Fusion gun turrets: %d, tons: %d",
				weapons.fusion, int(float32(2*weapons.fusion)+.9999)))
		fusion.Show()
	} else {
		fusion.Hide()
	}
	fusion.Refresh()
}

func (w *weaponsDetail) buildSand() {
	if weapons.sandcaster > 0 {
		sand.SetText(
			buildAmmoWeaponString("Triple Sandcaster turrets: %d, tons: %d, ammo tons: %d",
				weapons.sandcaster, int(float32(weapons.sandcaster)/2.0+.9999), int(float32(weapons.sandcaster)+.9999)))
		sand.Show()
	} else {
		sand.Hide()
	}
	sand.Refresh()
}

func (w *weaponsDetail) buildParticle() {
	if weapons.accelerator > 0 {
		particle.SetText(
			buildWeaponString("Particle Accelerator turrets: %d, tons: %d",
				weapons.accelerator, int(float32(3*weapons.accelerator)+0.9999)))
		particle.Show()
	} else {
		particle.Hide()
	}
	particle.Refresh()
}

func (w *weaponsDetail) countWeapons() int {
	return weapons.missile + weapons.beam + weapons.pulse + weapons.plasma + weapons.sandcaster +
		weapons.fusion + weapons.accelerator
}

func (w *weaponsDetail) buildWeapons() {
	w.buildMissile()
	w.buildBeam()
	w.buildPlasma()
	w.buildFusion()
	w.buildSand()
	w.buildParticle()
}

func (w *weaponsDetail) tons() int {
	return int(.9999 + (5.0*float32(weapons.missile) + 0.5*float32(weapons.sandcaster) +
		float32(weapons.beam) + float32(weapons.pulse) + 2.0*float32(weapons.fusion) +
		2.0*float32(weapons.plasma) + 5.0*float32(weapons.accelerator)))
}

func (w *weaponsDetail) cost() int {
	return int(float32(weapons.missile) + float32(weapons.sandcaster) +
		3.0*float32(weapons.beam) + 3.0*float32(weapons.pulse) + 4.0*float32(weapons.fusion) +
		4.0*float32(weapons.plasma) + 5.0*float32(weapons.accelerator))
}

func buildWeaponString(weaponDescription string, count int, tons int) string {
	if count > 0 {
		return fmt.Sprintf(weaponDescription, count, int(float32(tons)))
	}
	return ""
}

func buildAmmoWeaponString(weaponAmmoDescription string, count int, tons int, ammoTons int) string {
	if count > 0 {
		return fmt.Sprintf(weaponAmmoDescription, count, tons, int(.999+float32(ammoTons)))
	}
	return ""
}
