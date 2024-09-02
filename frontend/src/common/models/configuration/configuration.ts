interface ConstructorParams {
  defaults: Record<string, Record<string, any>>;
}

class Configuration {
  defaults: Record<string, Record<string, any>>;

  constructor({ defaults }: ConstructorParams) {
    this.defaults = defaults;
  }

  static fromJSON(obj: Record<string, any>) {
    return new Configuration(obj as ConstructorParams)
  }

  public get asJson(): Record<string, any> {
    return {
      defaults: this.defaults
    } as Record<string, any>;
  }
}

export default Configuration;
