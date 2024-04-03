def dfs(graph, start, end, path=[]):
    path = path + [start]
    if start == end:
        return [path]
    paths = []
    for node in graph[start]:
        if node not in path:
            new_paths = dfs(graph, node, end, path)
            for new_path in new_paths:
                paths.append(new_path)
    return paths

# Example graph represented as adjacency list
graph = {
    'A': [ 'C', 'G'],
    'C': ['A', 'F', 'G'],
    'D': [],
    'E': [],
    'F': ['C'],
    'B': ['G'],
    'G': ['C', 'A', 'B']
}

start_point = 'A'
end_point = 'B'

all_paths = dfs(graph, start_point, end_point)
print("All possible paths from", start_point, "to", end_point, ":", all_paths)
