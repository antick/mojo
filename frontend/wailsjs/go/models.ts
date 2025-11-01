export namespace configs {
	
	export class CombinedModType {
	    Replacements: Record<string, string>;
	
	    static createFrom(source: any = {}) {
	        return new CombinedModType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Replacements = source["Replacements"];
	    }
	}
	export class SubModType {
	    Enabled: boolean;
	    Replacements: Record<string, string>;
	
	    static createFrom(source: any = {}) {
	        return new SubModType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Enabled = source["Enabled"];
	        this.Replacements = source["Replacements"];
	    }
	}
	export class ModType {
	    CombinedMod: CombinedModType;
	    SubMods: Record<string, SubModType>;
	
	    static createFrom(source: any = {}) {
	        return new ModType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.CombinedMod = this.convertValues(source["CombinedMod"], CombinedModType);
	        this.SubMods = this.convertValues(source["SubMods"], SubModType, true);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

