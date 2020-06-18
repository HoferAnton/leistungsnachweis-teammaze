<h1 align="center">
    Team Maze
</h1>
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
to get development sources for gtk/openGL. <br>
Running:<br>
<code>sudo apt-get update && sudo ./install_dev_deps.sh</code><br>
in the root folder of the repo should be enough to get you started.<br>
Subsequently, you have to build the project with:<br>
<code>go build</code>,<br>
and run the main with:<br>
<code>go run main.go</code>.

### Usage
On top left corner of the display you have "Labyrinth" and a "View" tab.<br>
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
                If the auto rotation is disabled you can rotate the cube<br>
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
<code>NewLocation(x uint, y uint, z uint) Location</code>.<br>
<br>
The public "Labyrinth" interface contains:
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
                Takes a Location and returns all connected location of the provided location.<br>
                The connected locations is a subset of the neighbors.
            </li>
        </ul>
    </li>
    <li>
        <code>IsConnected(Location, Location) bool</code>
        <ul>
            <li>
                Proves if the two provided locations are connected.<br>
                Returns true if they are.
            </li>
        </ul>
    </li>
    <li>
        <code>Connect(Location, Location) bool</code>
        <ul>
            <li>
                Connects the two provided locations.<br>
                Returns true if the connection was successfully build.
            </li>
        </ul>
    </li>
    <li>
        <code>Disconnect(Location, Location) bool</code>
        <ul>
            <li>
                Removes the connection between the two locations.<br>
                Returns true if the connection was successfully removed.
            </li>
        </ul>
    </li>
    <li>
        <code>Compare(Labyrinth) bool</code>
        <ul>
            <li>
                Compares this to another Labyrinth.<br>
                Returns true if both have a equal amount of locations,<br>
                all locations are equal and all connections are equal.
            </li>
        </ul>
    </li>
    <li>
        <code>CheckLocation(Location) bool</code>
        <ul>
            <li>Checks if the location is part of the cube, spanned by (0, 0, 0) and MaxLocation.</li>
        </ul>
    </li>
</ul>
You can generate new Labyrinths with the function:<br>
<code>func NewLabyrinth(maxLoc Location) Labyrinth</code>.

### Generating
The public generator interface contains only one function:<br>
<code>GenerateLabyrinth(furthestPoint common.Location) (common.Labyrinth, []common.Pair)</code>.<br>
It generates a new labyrinth with the provided size. <br>
It returns the generated labyrinth as the first, and the steps of the generation as the second return value.<br>
You can get instances of the concrete BreathFirstGenerator with the function:<br>
<code>NewBreadthFirstGenerator() BreadthFirstGenerator</code>,<br>
and the DepthFirstGenerator with the function:<br>
<code>NewDepthFirstGenerator() DepthFirstGenerator</code>.

#### Depth first algorithm
<ol>
    <li>loc <- Start at a random location of the cube.</li>
    <li>list <- Get all neighbors of loc</li>
    <li>for elements of list that are not part of the labyrinth in a random order
        <ol>
            <li>remove the the wall between the loc and the element</li>
            <li>push loc on a stack</li>
            <li>loc <- element</li>
            <li>go to 2</li>
            <li>loc <- pop loc from stack</li>
        </ol>
    </li>
</ol>

#### Breath first algorithm
<ol>
    <li>loc <- Start at a random location of the cube.</li>
    <li>add loc to worklist</li>
    <li>while worklist has elements
        <li>elem <- a random element of the worklist</li>
        <li>remove the elem from worklist</li>
        <ol>
            <li>for all neighbors that are not part of the labyrinth as elem2
                <ol>
                    <li>remove the the wall between the elem and the elem2</li>
                    <li>add elem2 to the worklist</li>
                </ol>
            </li>
        </ol>
    </li>
</ol>

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
