# Team Maze
<p align="center">
  <a href="#requirements">Requirements</a> •
  <a href="#setup">Setup</a> •
  <a href="#usage">Usage</a> •
  <a href="#common">Common</a> •
  <a href="#generating">Generating</a> •
  <a href="#solving">Solving</a> •
  <a href="#visualizing">Visualizing</a> •
  <a href="#licensing">Licensing</a>
</p>

We are generating, solving and visualizing 3D mazes.

### Requirements
The Application requires a go version 1.13.3 or grater.

### Setup
To set up your dev environment you may have to install a few packages </br
to get development sources for gtk/openGL. </br>
Running:
<code>sudo apt-get update && sudo ./install_dev_deps.sh</code></br>
in the root folder of the repo should be enough to get you started.</br>
Subsequently, you have to build the project with:</br>
<code>go build</code>,</br>
and run the main with:</br>
<code>go run main.go</code>.

### Usage
On top left corner of the display you have "Labyrinth" and a "View" tab.</br>
In the tab "Labyrinth" you can select:
<ul>
    <li>"Generator" let you select a generation algorithm:
        <ul>
            <li>"DepthFirst" generates a Labyrinth with a randomized depth first algorithm</li>
            <li>"BreathFirst" generates a Labyrinth with a randomized breath first algorithm</li>
        </ul>
    </li>
    <li>"Solver" let you select a solving algorithm:
        <ul>
            <li>"Recursive" a recursive solving algorithm</li>
            <li>"Concurrent" a multithreading solving algorithm</li>
        </ul>
    </li>
    <li>"Generate Random" generates a new labyrinth with a random size</li>
</ul>
In the tab "View" you can select:
<ul>
    <li>"Dragging Enabled" let you enable/disable auto rotate
        <ul>
            <li>
                If the auto rotation is disabled you can rotate the cube</br>
                by clicking at the cube and move it with the mouse
            </li>
        </ul>
    </li>
    <li>"Show" let you select multiple view options:
        <ul>
            <li>"Solver Path" shows the solution (0, 0, 0) to the opposite corner</li>
            <li>"Solver Algorithm" shows the steps of the solving algorithm</li>
            <li>"Generator Algorithm" shows the steps of the generation algorithm</li>
        </ul>
    </li>
</ul>

### Common
The public "Location" interface contains:
<ul>
    <li>
        <code>As3DCoordinates() (uint, uint, uint)</code>
        <ul>
            <li>
                It is a getter for the tuple form of the location.
                It has the form (x, y, z).
            </li>
        </ul>
    </li>
    <li>
        <code>Compare(Location) bool</code>
        <ul>
            <li>
                It compares this location with another location.
                Both are equal if the x, y and z parts are equal.
             </li>
        </ul>
    </li>
</ul>
You can generate new locations with the function:<br/>
<code>NewLocation(x uint, y uint, z uint) Location</code>.</br>
</br>
The public "Labyrint" interface contains:
<ul>
    <li>
        <code>GetMaxLocation() Location</code>
        <ul>
            <li>It is a getter for the furthest point in the labyrinth form the root (0, 0, 0).</li>
        </ul>
    </li>
    <li>
        <code>GetNeighbors(Location) []Location</code>
        <ul>
            <li>Takes a Location and returns all neighbors of this location.</li>
        </ul>
    </li>
    <li>
        <code>GetConnected(Location) []Location</code>
        <ul>
            <li>
                Takes a Location and returns all connected location of the provided location.</br>
                The connected locations is a subset of the neighbors.
            </li>
        </ul>
    </li>
    <li>
        <code>IsConnected(Location, Location) bool</code>
        <ul>
            <li>
                Proves if the two provided locations are connected.</br>
                Returns true if they are.
            </li>
        </ul>
    </li>
    <li>
        <code>Connect(Location, Location) bool</code>
        <ul>
            <li>
                Connects the two provided locations.</br>
                Returns true if the connection was successfully build.
            </li>
        </ul>
    </li>
    <li>
        <code>Disconnect(Location, Location) bool</code>
        <ul>
            <li>
                Removes the connection between the two locations.</br>
                Returns true if the connection was successfully removed.
            </li>
        </ul>
    </li>
    <li>
        <code>Compare(Labyrinth) bool</code>
        <ul>
            <li>
                Compares this to another Labyrinth.</br>
                Returns true if both have a equal amount of locations,</br>
                all locations are equal and all connections are equal.
            </li>
        </ul>
    </li>
    <li>
        <code>CheckLocation(Location) bool</code>
        <ul>
            <li>Checks if the location is </li>
        </ul>
    </li>
    <li>
        <code>getNode(Location) Node</code>
        <ul>
            <li></li>
        </ul>
    </li>
</ul>

### Generating

#### Depth first algorithm

#### Breath first algorithm


### Solving
TODO: describe solving interface
#### Recursive solver algorithm
TODO: describe recursive solver algorithm
#### Concurrent solver algorithm
TODO: describe concurrent solver algorithm

### Visualizing

TODO: describe visualizing

### Licensing
TODO: select a license (e.g. MIT?)
