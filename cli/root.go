/*
 * moly - a game engine inspired by better games
 * Copyright (C) 2022 Michael D Henderson
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// cmdRoot represents the base command when called without any subcommands
var cmdRoot = &cobra.Command{
	Short:   "Moly: A game engine",
	Long:    `moly is the engine inspired by better games.`,
	Version: "0.0.1",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		started := time.Now()
		if argsRoot.add_flag {
			argsRoot.immediate = false
		}
		if argsRoot.eat_flag {
			argsRoot.immediate = false
		}
		if argsRoot.run_flag {
			argsRoot.immediate = false
		}

		log.Printf("%-20s == %q\n", "lib-dir", argsRoot.libdir)

		if argsRoot.combat_test_flag {
			if argsRoot.libdir == "" {
				return fmt.Errorf("missing lib-dir argument")
			}
		}

		if argsRoot.time_self {
			elapsed := time.Now().Sub(started)
			log.Printf("elapsed time: %v\n", elapsed)
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the root Command.
func Execute() error {
	//fprintf(stderr, "usage: moly [options]\n");
	//fprintf(stderr, "  -a        Add new players mode\n");
	//fprintf(stderr, "  -e        Eat orders from libdir/spool\n");
	//fprintf(stderr, "  -f        Don't buffer files for debugging\n");
	//fprintf(stderr, "  -i        Immediate mode\n");
	//fprintf(stderr, "  -l dir    Specify libdir, default ./lib\n");
	//fprintf(stderr, "  -p        Don't make data files pretty\n");
	//fprintf(stderr, "  -r        Run a turn\n");
	//fprintf(stderr, "  -t        Test ilist code\n");
	//fprintf(stderr, "  -x        Inhibit adding players during turn.\n");
	//fprintf(stderr, "  -A        Charge player accounts\n");
	//fprintf(stderr, "  -L        Generate lore dictionary.\n");
	//fprintf(stderr, "  -M        Mail reports\n");
	//fprintf(stderr, "  -R        Test the random number generator\n");
	//fprintf(stderr, "  -S        Save the database at completion\n");
	//fprintf(stderr, "  -T        Print timing info\n");
	//fprintf(stderr, "  -X        Combat test\n");

	cmdRoot.PersistentFlags().StringVar(&argsRoot.libdir, "lib-dir", "", "set lib path")
	cmdRoot.PersistentFlags().BoolVar(&argsRoot.save_flag, "save-db", false, "set save-db-flag")
	cmdRoot.PersistentFlags().BoolVar(&argsRoot.time_self, "time", false, "time commands")

	cmdRoot.Flags().BoolVar(&argsRoot.add_flag, "a", false, "set add-flag")
	cmdRoot.Flags().BoolVar(&argsRoot.eat_flag, "e", false, "set eat-flag")
	cmdRoot.Flags().BoolVar(&argsRoot.flush_always, "f", false, "set flush-always-flag")
	cmdRoot.Flags().BoolVar(&argsRoot.immed_after, "i", false, "set immed-after-flag")
	cmdRoot.Flags().BoolVar(&argsRoot.map_flag, "m", false, "set map-flag")
	cmdRoot.Flags().BoolVar(&argsRoot.pretty_data_files, "p", false, "set pretty-data-files-flag")
	cmdRoot.Flags().BoolVar(&argsRoot.art_flag, "q", false, "set test-artifacts-flag")
	cmdRoot.Flags().BoolVar(&argsRoot.run_flag, "r", false, "set run-turn-flag")
	cmdRoot.Flags().BoolVar(&argsRoot.inhibit_add_flag, "x", false, "set inhibit-add-flag")
	cmdRoot.Flags().BoolVar(&argsRoot.acct_flag, "A", false, "set acct-flag")
	cmdRoot.Flags().BoolVar(&argsRoot.unspool_first_flag, "E", false, "set unspool-first-flag")
	cmdRoot.Flags().BoolVar(&argsRoot.lore_flag, "L", false, "set lore-flag")
	cmdRoot.Flags().BoolVar(&argsRoot.mail_now, "M", false, "set mail-now-flag")

	cmdRoot.Flags().BoolVar(&argsRoot.combat_test_flag, "test-combat", false, "set combat-test-flag")
	cmdRoot.Flags().BoolVar(&argsRoot.testJsonLoad, "test-json-load", false, "test load from json store")
	cmdRoot.Flags().BoolVar(&argsRoot.test_lists_flag, "t", false, "set test-lists-flag")
	cmdRoot.Flags().BoolVar(&argsRoot.test_prng_flag, "R", false, "set test-prng-flag")

	cmdRoot.AddCommand(cmdGenerate)

	cmdGenerate.AddCommand(cmdGenerateMap)
	// inputs
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.mapFileName, "map-data", "worldmap.txt", "map data to import")
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.cityFileName, "city-data", "cities.json", "city name data to import")
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.landFileName, "land-data", "lands.json", "land data to import")
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.regionFileName, "region-data", "regions.json", "region data to import")
	cmdGenerateMap.Flags().BoolVar(&argsGenerateMap.noColor, "no-color", argsGenerateMap.noColor, "do not color map")
	cmdGenerateMap.Flags().BoolVar(&argsGenerateMap.noScale, "no-scale", argsGenerateMap.noScale, "do not scale map")
	cmdGenerateMap.Flags().BoolVar(&argsGenerateMap.createPNG, "create-png", argsGenerateMap.createPNG, "create a png of the map")
	cmdGenerateMap.Flags().BoolVar(&argsGenerateMap.createSVG, "create-svg", argsGenerateMap.createSVG, "create an svg of the map")
	// outputs
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.continentFileName, "continent-data", "continents.json", "continent data to export")
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.locationFileName, "location-data", "locations.json", "location data to export")
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.gateFileName, "gate-data", "gates.json", "gate data to export")
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.roadFileName, "road-data", "roads.json", "road data to export")
	cmdGenerateMap.Flags().StringVar(&argsGenerateMap.seedFileName, "seed-data", "randseed.json", "random seed data to export")

	cmdRoot.AddCommand(cmdServe)
	cmdServe.Flags().StringVar(&argsServe.host, "host", argsServe.host, "host to bind to")
	cmdServe.Flags().BoolVar(&argsServe.logRequests, "log-requests", argsServe.logRequests, "log all requests")
	cmdServe.Flags().StringVar(&argsServe.port, "port", argsServe.port, "port to listen on")
	cmdServe.Flags().StringVar(&argsServe.public, "public", argsServe.public, "path to public files")
	if err := cmdServe.MarkFlagRequired("public"); err != nil {
		panic(err)
	}

	return cmdRoot.Execute()
}

var argsRoot struct {
	acct_flag          bool
	add_flag           bool
	art_flag           bool
	combat_test_flag   bool
	eat_flag           bool
	flush_always       bool
	immed_after        bool
	immediate          bool
	inhibit_add_flag   bool
	libdir             string
	lore_flag          bool
	mail_now           bool
	map_flag           bool
	map_test_flag      bool
	pretty_data_files  bool
	run_flag           bool
	save_flag          bool
	test_prng_flag     bool
	test_lists_flag    bool
	testJsonLoad       bool
	time_self          bool
	unspool_first_flag bool
}
