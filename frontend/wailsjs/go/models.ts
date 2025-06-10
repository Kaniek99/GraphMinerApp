export namespace main {
	
	export class InputFile {
	    Extension: string;
	    Path: string;
	    Valid: boolean;
	    Parser: any;
	
	    static createFrom(source: any = {}) {
	        return new InputFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Extension = source["Extension"];
	        this.Path = source["Path"];
	        this.Valid = source["Valid"];
	        this.Parser = source["Parser"];
	    }
	}

}

