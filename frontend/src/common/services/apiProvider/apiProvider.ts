class APIProvider {
// SINGLETON ---------------------
  private constructor() { }
  private _instance: APIProvider|null = null

  public get instance() : APIProvider {
    if (this._instance == null) {
      this._instance = new APIProvider();
    }

    return this._instance;
  }

// METHODS ---------------------
  async setMatrix(module: string, opts: Record<string, any>) : Promise<string|null> {
    try {
      const result = await fetch(`/api/${module}`, {
        method:"POST",
        body: JSON.stringify(opts),
      });
      const json = await result.json();
      if (json.success) return null;
      return json.error as string;
    } catch (e) {
      return `Something went wrong: ${e}`;
    }
  }
}

export default APIProvider;

